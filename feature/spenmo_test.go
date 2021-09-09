package feature_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/cucumber/godog"
	"github.com/indrasaputra/hashids"
	"github.com/joho/godotenv"
)

var (
	ctx     = context.Background()
	client  = http.DefaultClient
	cardURL = "http://localhost:8081/v1/users/cards"

	httpStatus int
	httpBody   []byte
)

// Card defines logical data for user's card.
type Card struct {
	ID           string  `json:"id"`
	UserID       string  `json:"userId"`
	WalletID     string  `json:"walletId"`
	LimitDaily   float64 `json:"limitDaily"`
	LimitMonthly float64 `json:"limitMonthly"`
}

type GetSingleResponse struct {
	Card *Card `json:"card"`
}

type GetAllResponse struct {
	Cards []*Card `json:"cards"`
}

func TestMain(_ *testing.M) {
	status := godog.TestSuite{
		Name:                "spenmo v1alpha1",
		ScenarioInitializer: InitializeScenario,
	}.Run()

	os.Exit(status)
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	_ = godotenv.Load()
	url := os.Getenv("SERVER_URL")
	if url != "" {
		cardURL = url
	}

	ctx.Step(`^I create card owned by user (\d+) with body$`, iCreateCardOwnedByUserWithBody)
	ctx.Step(`^response must match json$`, responseMustMatchJson)
	ctx.Step(`^response status code must be (\d+)$`, responseStatusCodeMustBe)
	ctx.Step(`^the card owned by user (\d+) is empty$`, theCardOwnedByUserIsEmpty)
	ctx.Step(`^I delete card with id (\d+) owned by user (\d+)$`, iDeleteCardWithIdOwnedByUser)
	ctx.Step(`^I delete card with index (\d+) owned by user (\d+)$`, iDeleteCardWithIndexOwnedByUser)
	ctx.Step(`^there are cards owned by user (\d+) with body$`, thereAreCardsOwnedByUserWithBody)
	ctx.Step(`^I get all cards owned by user (\d+)$`, iGetAllCardsOwnedByUser)
	ctx.Step(`^I get single card with id (\d+) owned by user (\d+)$`, iGetSingleCardWithIdOwnedByUser)
	ctx.Step(`^I get single card with index (\d+) owned by user (\d+)$`, iGetSingleCardWithIndexOwnedByUser)
	ctx.Step(`^response must be single card$`, responseMustBeSingleCard)
	ctx.Step(`^number of cards retrieved must be (\d+)$`, numberOfCardsRetrievedMustBe)
	ctx.Step(`^I update card with id (\d+) owned by user (\d+) with body$`, iUpdateCardWithIdOwnedByUserWithBody)
	ctx.Step(`^I update card with index (\d+) owned by user (\d+) with body$`, iUpdateCardWithIndexOwnedByUserWithBody)
}

func iCreateCardOwnedByUserWithBody(userID int, requests *godog.Table) error {
	for _, row := range requests.Rows {
		body := strings.NewReader(row.Cells[0].Value)
		if err := callEndpoint(http.MethodPost, cardURL, userID, body); err != nil {
			return err
		}
	}
	return nil
}

func responseMustMatchJson(want *godog.DocString) error {
	return deepCompareJSON([]byte(want.Content), httpBody)
}

func responseStatusCodeMustBe(code int) error {
	if httpStatus != code {
		return fmt.Errorf("expected HTTP status code %d, but got %d", code, httpStatus)
	}
	return nil
}

func theCardOwnedByUserIsEmpty(userID int) error {
	return deleteAll(userID)
}

func iDeleteCardWithIdOwnedByUser(cardID, userID int) error {
	cid := hashids.ID(int64(cardID)).EncodeString()
	return callEndpoint(http.MethodDelete, fmt.Sprintf("%s/%s", cardURL, cid), userID, nil)
}

func iDeleteCardWithIndexOwnedByUser(cardIndex, userID int) error {
	cards, err := getAllCards(userID)
	if err != nil {
		return fmt.Errorf("get all cards error: %v", err)
	}

	for i := range cards {
		if i == cardIndex {
			id, _ := hashids.DecodeHash([]byte(cards[i].ID))
			return iDeleteCardWithIdOwnedByUser(int(id), userID)
		}
	}
	return fmt.Errorf("no card can be deleted")
}

func thereAreCardsOwnedByUserWithBody(userID int, requests *godog.Table) error {
	return iCreateCardOwnedByUserWithBody(userID, requests)
}

func iGetAllCardsOwnedByUser(userID int) error {
	return callEndpoint(http.MethodGet, cardURL, userID, nil)
}

func iGetSingleCardWithIdOwnedByUser(cardID, userID int) error {
	cid := hashids.ID(int64(cardID)).EncodeString()
	return callEndpoint(http.MethodGet, fmt.Sprintf("%s/%s", cardURL, cid), userID, nil)
}

func iGetSingleCardWithIndexOwnedByUser(cardIndex, userID int) error {
	cards, err := getAllCards(userID)
	if err != nil {
		return fmt.Errorf("get all cards error: %v", err)
	}

	for i := range cards {
		if i == cardIndex {
			id, _ := hashids.DecodeHash([]byte(cards[i].ID))
			return iGetSingleCardWithIdOwnedByUser(int(id), userID)
		}
	}
	return fmt.Errorf("no card can be retrieved")
}

func responseMustBeSingleCard() error {
	var actual *GetSingleResponse

	err := json.Unmarshal(httpBody, &actual)
	if err != nil {
		return err
	}

	if actual != nil {
		return nil
	}
	return fmt.Errorf("response is not single card")
}

func numberOfCardsRetrievedMustBe(count int) error {
	var actual GetAllResponse

	err := json.Unmarshal(httpBody, &actual)
	if err != nil {
		return err
	}

	if len(actual.Cards) == count {
		return nil
	}
	return fmt.Errorf("expected number of cards %d, but got %d", count, len(actual.Cards))
}

func iUpdateCardWithIdOwnedByUserWithBody(cardID, userID int, requests *godog.Table) error {
	for _, row := range requests.Rows {
		body := strings.NewReader(row.Cells[0].Value)
		cid := hashids.ID(int64(cardID)).EncodeString()
		if err := callEndpoint(http.MethodPut, fmt.Sprintf("%s/%s", cardURL, cid), userID, body); err != nil {
			return err
		}
	}
	return nil
}

func iUpdateCardWithIndexOwnedByUserWithBody(cardIndex, userID int, requests *godog.Table) error {
	cards, err := getAllCards(userID)
	if err != nil {
		return fmt.Errorf("get all cards error: %v", err)
	}

	for i := range cards {
		if i == cardIndex {
			id, _ := hashids.DecodeHash([]byte(cards[i].ID))
			return iUpdateCardWithIdOwnedByUserWithBody(int(id), userID, requests)
		}
	}
	return fmt.Errorf("no card can be updated")
}

func getAllCards(userID int) ([]*Card, error) {
	if err := callEndpoint(http.MethodGet, cardURL, userID, nil); err != nil {
		return nil, err
	}

	var resp GetAllResponse
	if err := json.Unmarshal(httpBody, &resp); err != nil {
		return nil, err
	}

	return resp.Cards, nil
}

func deleteAll(userID int) error {
	cards, err := getAllCards(userID)
	if err != nil {
		return err
	}

	for _, card := range cards {
		if err = callEndpoint(http.MethodDelete, fmt.Sprintf("%s/%s", cardURL, card.ID), userID, nil); err != nil {
			return err
		}
	}
	return nil
}

func callEndpoint(method, url string, userID int, body io.Reader) error {
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", fmt.Sprintf("%d", userID))
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	httpStatus = resp.StatusCode
	httpBody, err = ioutil.ReadAll(resp.Body)
	return err
}

func deepCompareJSON(want, have []byte) error {
	var expected interface{}
	var actual interface{}

	err := json.Unmarshal(want, &expected)
	if err != nil {
		return err
	}
	err = json.Unmarshal(have, &actual)
	if err != nil {
		return err
	}

	if !reflect.DeepEqual(expected, actual) {
		return fmt.Errorf("expected JSON does not match actual, %v vs. %v", expected, actual)
	}
	return nil
}
