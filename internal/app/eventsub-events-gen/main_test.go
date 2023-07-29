package main

import (
	"strings"
	"testing"

	"github.com/vpetrigo/go-twitch-ws/internal/pkg/crawler"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

const (
	charityEventTable = `<table>
  <thead>
    <tr>
      <th>Field</th>
      <th>Type</th>
      <th>Description</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td><code class="highlighter-rouge">id</code></td>
      <td>String</td>
      <td>An ID that identifies the charity campaign.</td>
    </tr>
    <tr>
      <td><code class="highlighter-rouge">broadcaster_id</code></td>
      <td>String</td>
      <td>An ID that identifies the broadcaster that’s running the campaign.</td>
    </tr>
    <tr>
      <td><code class="highlighter-rouge">broadcaster_login</code></td>
      <td>String</td>
      <td>The broadcaster’s login name.</td>
    </tr>
    <tr>
      <td><code class="highlighter-rouge">broadcaster_name</code></td>
      <td>String</td>
      <td>The broadcaster’s display name.</td>
    </tr>
    <tr>
      <td><code class="highlighter-rouge">charity_name</code></td>
      <td>String</td>
      <td>The charity’s name.</td>
    </tr>
    <tr>
      <td><code class="highlighter-rouge">charity_description</code></td>
      <td>String</td>
      <td>A description of the charity.</td>
    </tr>
    <tr>
      <td><code class="highlighter-rouge">charity_logo</code></td>
      <td>String</td>
      <td>A URL to an image of the charity’s logo. The image’s type is PNG and its size is 100px X 100px.</td>
    </tr>
    <tr>
      <td><code class="highlighter-rouge">charity_website</code></td>
      <td>String</td>
      <td>A URL to the charity’s website.</td>
    </tr>
    <tr>
      <td><code class="highlighter-rouge">current_amount</code></td>
      <td>Object</td>
      <td>An object that contains the current amount of donations that the campaign has received.</td>
    </tr>
    <tr>
      <td>&nbsp;&nbsp;&nbsp;<code class="highlighter-rouge">value</code></td>
      <td>Integer</td>
      <td>The monetary amount. The amount is specified in the currency’s minor unit. For example, the minor units for USD is cents, so if the amount is $5.50 USD, <code class="highlighter-rouge">value</code> is set to 550.</td>
    </tr>
    <tr>
      <td>&nbsp;&nbsp;&nbsp;<code class="highlighter-rouge">decimal_places</code></td>
      <td>Integer</td>
      <td>The number of decimal places used by the currency. For example, USD uses two decimal places. Use this number to translate <code class="highlighter-rouge">value</code> from minor units to major units by using the formula:<br><br><code class="highlighter-rouge">value / 10^decimal_places</code></td>
    </tr>
    <tr>
      <td>&nbsp;&nbsp;&nbsp;<code class="highlighter-rouge">currency</code></td>
      <td>String</td>
      <td>The ISO-4217 three-letter currency code that identifies the type of currency in <code class="highlighter-rouge">value</code>.</td>
    </tr>
    <tr>
      <td><code class="highlighter-rouge">target_amount</code></td>
      <td>Object</td>
      <td>An object that contains the campaign’s target fundraising goal.</td>
    </tr>
    <tr>
      <td>&nbsp;&nbsp;&nbsp;<code class="highlighter-rouge">value</code></td>
      <td>Integer</td>
      <td>The monetary amount. The amount is specified in the currency’s minor unit. For example, the minor units for USD is cents, so if the amount is $5.50 USD, <code class="highlighter-rouge">value</code> is set to 550.</td>
    </tr>
    <tr>
      <td>&nbsp;&nbsp;&nbsp;<code class="highlighter-rouge">decimal_places</code></td>
      <td>Integer</td>
      <td>The number of decimal places used by the currency. For example, USD uses two decimal places. Use this number to translate <code class="highlighter-rouge">value</code> from minor units to major units by using the formula:<br><br><code class="highlighter-rouge">value / 10^decimal_places</code></td>
    </tr>
    <tr>
      <td>&nbsp;&nbsp;&nbsp;<code class="highlighter-rouge">currency</code></td>
      <td>String</td>
      <td>The ISO-4217 three-letter currency code that identifies the type of currency in <code class="highlighter-rouge">value</code>.</td>
    </tr>
    <tr>
      <td><code class="highlighter-rouge">started_at</code></td>
      <td>String</td>
      <td>The UTC timestamp (in RFC3339 format) of when the broadcaster started the campaign.</td>
    </tr>
  </tbody>
</table>`
	standardEventTable = `<table>
  <thead>
    <tr>
      <th>Name</th>
      <th>Type</th>
      <th>Description</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td><code class="highlighter-rouge">user_id</code></td>
      <td>string</td>
      <td>The user ID of the user who sent a resubscription chat message.</td>
    </tr>
    <tr>
      <td><code class="highlighter-rouge">user_login</code></td>
      <td>string</td>
      <td>The user login of the user who sent a resubscription chat message.</td>
    </tr>
    <tr>
      <td><code class="highlighter-rouge">user_name</code></td>
      <td>string</td>
      <td>The user display name of the user who a resubscription chat message.</td>
    </tr>
    <tr>
      <td><code class="highlighter-rouge">broadcaster_user_id</code></td>
      <td>string</td>
      <td>The broadcaster user ID.</td>
    </tr>
    <tr>
      <td><code class="highlighter-rouge">broadcaster_user_login</code></td>
      <td>string</td>
      <td>The broadcaster login.</td>
    </tr>
    <tr>
      <td><code class="highlighter-rouge">broadcaster_user_name</code></td>
      <td>string</td>
      <td>The broadcaster display name.</td>
    </tr>
    <tr>
      <td><code class="highlighter-rouge">tier</code></td>
      <td>string</td>
      <td>The tier of the user’s subscription.</td>
    </tr>
    <tr>
      <td><code class="highlighter-rouge">message</code></td>
      <td><a href="#message">message</a></td>
      <td>An object that contains the resubscription message and emote information needed to recreate the message.</td>
    </tr>
    <tr>
      <td><code class="highlighter-rouge">cumulative_months</code></td>
      <td>integer</td>
      <td>The total number of months the user has been subscribed to the channel.</td>
    </tr>
    <tr>
      <td><code class="highlighter-rouge">streak_months</code></td>
      <td>integer</td>
      <td>The number of consecutive months the user’s current subscription has been active. This value is <code class="highlighter-rouge">null</code> if the user has opted out of sharing this information.</td>
    </tr>
    <tr>
      <td><code class="highlighter-rouge">duration_months</code></td>
      <td>integer</td>
      <td>The month duration of the subscription.</td>
    </tr>
  </tbody>
</table>`
	dummyTable = `<table>
  <thead>
    <tr>
      <th>Foo</th>
      <th>Bar</th>
      <th>Baz</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td><code class="highlighter-rouge">id</code></td>
      <td>String</td>
      <td>An ID that identifies the charity campaign.</td>
    </tr>
    <tr>
      <td><code class="highlighter-rouge">broadcaster_id</code></td>
      <td>String</td>
      <td>An ID that identifies the broadcaster that’s running the campaign.</td>
    </tr>
    <tr>
      <td><code class="highlighter-rouge">broadcaster_login</code></td>
      <td>String</td>
      <td>The broadcaster’s login name.</td>
    </tr>
    <tr>
      <td><code class="highlighter-rouge">broadcaster_name</code></td>
      <td>String</td>
      <td>The broadcaster’s display name.</td>
    </tr>
    <tr>
      <td><code class="highlighter-rouge">charity_name</code></td>
      <td>String</td>
      <td>The charity’s name.</td>
    </tr>
    <tr>
      <td><code class="highlighter-rouge">charity_description</code></td>
      <td>String</td>
      <td>A description of the charity.</td>
    </tr>
    <tr>
      <td><code class="highlighter-rouge">charity_logo</code></td>
      <td>String</td>
      <td>A URL to an image of the charity’s logo. The image’s type is PNG and its size is 100px X 100px.</td>
    </tr>
    <tr>
      <td><code class="highlighter-rouge">charity_website</code></td>
      <td>String</td>
      <td>A URL to the charity’s website.</td>
    </tr>
    <tr>
      <td><code class="highlighter-rouge">current_amount</code></td>
      <td>Object</td>
      <td>An object that contains the current amount of donations that the campaign has received.</td>
    </tr>
    <tr>
      <td>&nbsp;&nbsp;&nbsp;<code class="highlighter-rouge">value</code></td>
      <td>Integer</td>
      <td>The monetary amount. The amount is specified in the currency’s minor unit. For example, the minor units for USD is cents, so if the amount is $5.50 USD, <code class="highlighter-rouge">value</code> is set to 550.</td>
    </tr>
    <tr>
      <td>&nbsp;&nbsp;&nbsp;<code class="highlighter-rouge">decimal_places</code></td>
      <td>Integer</td>
      <td>The number of decimal places used by the currency. For example, USD uses two decimal places. Use this number to translate <code class="highlighter-rouge">value</code> from minor units to major units by using the formula:<br><br><code class="highlighter-rouge">value / 10^decimal_places</code></td>
    </tr>
    <tr>
      <td>&nbsp;&nbsp;&nbsp;<code class="highlighter-rouge">currency</code></td>
      <td>String</td>
      <td>The ISO-4217 three-letter currency code that identifies the type of currency in <code class="highlighter-rouge">value</code>.</td>
    </tr>
    <tr>
      <td><code class="highlighter-rouge">target_amount</code></td>
      <td>Object</td>
      <td>An object that contains the campaign’s target fundraising goal.</td>
    </tr>
    <tr>
      <td>&nbsp;&nbsp;&nbsp;<code class="highlighter-rouge">value</code></td>
      <td>Integer</td>
      <td>The monetary amount. The amount is specified in the currency’s minor unit. For example, the minor units for USD is cents, so if the amount is $5.50 USD, <code class="highlighter-rouge">value</code> is set to 550.</td>
    </tr>
    <tr>
      <td>&nbsp;&nbsp;&nbsp;<code class="highlighter-rouge">decimal_places</code></td>
      <td>Integer</td>
      <td>The number of decimal places used by the currency. For example, USD uses two decimal places. Use this number to translate <code class="highlighter-rouge">value</code> from minor units to major units by using the formula:<br><br><code class="highlighter-rouge">value / 10^decimal_places</code></td>
    </tr>
    <tr>
      <td>&nbsp;&nbsp;&nbsp;<code class="highlighter-rouge">currency</code></td>
      <td>String</td>
      <td>The ISO-4217 three-letter currency code that identifies the type of currency in <code class="highlighter-rouge">value</code>.</td>
    </tr>
    <tr>
      <td><code class="highlighter-rouge">started_at</code></td>
      <td>String</td>
      <td>The UTC timestamp (in RFC3339 format) of when the broadcaster started the campaign.</td>
    </tr>
  </tbody>
</table>`
)

type testFixture struct {
	tableHTML    string
	expected     bool
	validationFn func(node *html.Node) bool
}

func Test_standardEventTableValidator(t *testing.T) {
	fixture := []testFixture{
		{
			tableHTML: charityEventTable,
			expected:  false,
		},
		{
			tableHTML: standardEventTable,
			expected:  true,
		},
		{
			tableHTML: dummyTable,
			expected:  false,
		},
	}

	testWithFixture(fixture, t, standardEventTableValidator)
}

func Test_charityEventTableValidator(t *testing.T) {
	fixture := []testFixture{
		{
			tableHTML: charityEventTable,
			expected:  true,
		},
		{
			tableHTML: standardEventTable,
			expected:  false,
		},
		{
			tableHTML: dummyTable,
			expected:  false,
		},
	}

	testWithFixture(fixture, t, charityEventTableValidator)
}

func testWithFixture(fixture []testFixture, t *testing.T, validationFn func(node *html.Node) bool) {
	for _, v := range fixture {
		n, _ := html.ParseFragment(strings.NewReader(v.tableHTML), &html.Node{
			Type:     html.ElementNode,
			Data:     "body",
			DataAtom: atom.Body,
		})

		var table *html.Node

		for _, e := range n {
			if crawler.IsElementNode(e) && e.Data == "table" {
				table = e
				break
			}
		}

		thead := skipToTableRow(table)

		if v.expected != validationFn(thead) {
			t.Fatalf("validation failed for: %s", v.tableHTML)
		}
	}
}

func skipToTableRow(node *html.Node) *html.Node {
	var tableHead *html.Node

	for it := node.FirstChild; it != nil; it = it.NextSibling {
		if crawler.IsElementNode(it) && it.Data == "thead" {
			tableHead = it
			break
		}
	}

	for it := tableHead.FirstChild; it != nil; it = it.NextSibling {
		if crawler.IsElementNode(it) && it.Data == "tr" {
			return it
		}
	}

	return nil
}
