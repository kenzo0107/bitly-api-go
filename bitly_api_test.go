package bitly_api

import (
	"os"
	"testing"
)

var linkUrl string = "http://bitly.com/Xlq5ZH"
var testUrl string = "http://bitly.com/3hQYj"
var longUrl string = "http://google.com"

func getConnection(t *testing.T) *Connection {

	token := "BITLY_ACCESS_TOKEN"

	BITLY_ACCESS_TOKEN := os.Getenv(token)
	if BITLY_ACCESS_TOKEN == "" {
		t.Fatalf(token + " not found")
		return nil
	}
	accessToken := BITLY_ACCESS_TOKEN
  return NewConnection(accessToken, "")
}

func TestApi(t *testing.T) {
	bitly := getConnection(t)
	testUrl := "http://google.com/"
	_, err := bitly.Shorten(testUrl)
	if err != nil {
		t.Fatalf("bitly Shorten returned an err %s", err)
	}
}

func TestExpand(t *testing.T) {
	bitly := getConnection(t)
	
  data, err := bitly.Expand("test1_random_fjslfjieljfklsjflkas")
	if err != nil {
		t.Fatalf("bitly Expand returned an error %s", err)
	}
	if data["error"] != "NOT_FOUND" {
		t.Fatalf("bitly Expand did not return NOT_FOUND", err)
	}
}

func TestClicks(t *testing.T) {
	bitly := getConnection(t)
  
  data, err := bitly.Clicks("test1_random_fjslfjieljfklsjflkas")
	if err != nil {
		t.Fatalf("bitly clicks returned an error %s", err)
	}
	if data["error"] != "NOT_FOUND" {
		t.Fatalf("bitly ClicksHash did not return NOT_FOUND", err)
	}

	data, err = bitly.ClicksByDay("test1_random_fjslfjieljfklsjflkas")
	if err != nil {
		t.Fatalf("bitly clicks returned an error %s", err)
	}
	if data["error"] != "NOT_FOUND" {
		t.Fatalf("bitly ClicksHash did not return NOT_FOUND", err)
	}

	data, err = bitly.ClicksByMinute("test1_random_fjslfjieljfklsjflkas")
	if err != nil {
		t.Fatalf("bitly clicks returned an error %s", err)
	}
	if data["error"] != "NOT_FOUND" {
		t.Fatalf("bitly ClicksHash did not return NOT_FOUND", err)
	}
	data, err = bitly.Clicks("http://bit.ly/3hQYj")
	if err != nil {
		t.Fatalf("bitly clicks returned an error %s", err)
	}
	if data["error"] == "NOT_FOUND" {
		t.Fatalf("bitly ClicksHash did not return NOT_FOUND", err)
	}

	data, err = bitly.ClicksByDay("http://bit.ly/3hQYj")
	if err != nil {
		t.Fatalf("bitly clicks returned an error %s", err)
	}
	if data["error"] == "NOT_FOUND" {
		t.Fatalf("bitly ClicksHash did not return NOT_FOUND", err)
	}

	data, err = bitly.ClicksByMinute("http://bit.ly/3hQYj")
	if err != nil {
		t.Fatalf("bitly clicks returned an error %s", err)
	}
	if data["error"] == "NOT_FOUND" {
		t.Fatalf("bitly ClicksHash did not return NOT_FOUND %s", err)
	}
}

func TestInfo(t *testing.T) {
	bitly := getConnection(t)

  _, err := bitly.Info(testUrl)
	if err != nil {
		t.Fatalf("bitly Info returned an error %s", err)
	}
}

func TestLinkEncodersCount(t *testing.T) {
	bitly := getConnection(t)

  _, err := bitly.LinkEncodersCount(testUrl)
	if err != nil {
		t.Fatalf("bitly links_encoders_count returned an error %s", err)
	}
}

func TestUserLink(t *testing.T) {
  bitly := getConnection(t)

  _, err := bitly.UserLinkLookup(testUrl)
	if err != nil {
	  t.Fatalf("bitly UserLinkLookup returned an error %s", err)
	}
  
  _, err = bitly.UserLinkSave(longUrl, UserLink{private:true})
  if err != nil && err.Error() != "LINK_ALREADY_EXISTS" {
    t.Fatalf("bitly user/link_save returned an expected result %s", err)
  }

  _, err = bitly.UserLinkEdit(linkUrl, "title", UserLink{title:"New Title"})
  if err != nil {
    t.Fatalf("bitly user/link_edit returned an expected result %s", err)
  }

  _, err = bitly.UserLinkHistory(UserLinkHistory{archived:"on"})
  if err != nil {
    t.Fatalf("bitly user/link_history returned an expected result %s", err)
  }
}

func TestLinkMetrics(t *testing.T) {
  bitly := getConnection(t)
  _, err := bitly.LinkClicks(testUrl, Metrics{limit:1})
  if err != nil {
    t.Fatalf("bitly link/clicks returned an error %s", err)
  }

  _, err = bitly.LinkReferrersByDomain(testUrl, Metrics{limit:1})
  if err != nil {
    t.Fatalf("bitly link/referrers_by_domain returned an error %s", err)
  }
  
  _, err = bitly.LinkReferrers(testUrl, Metrics{limit:1})
  if err != nil {
    t.Fatalf("bitly link/referrers returned an error %s", err)
  }

  _, err = bitly.LinkShares(testUrl, Metrics{limit:1})
  if err != nil {
    t.Fatalf("bitly link/shares returned an error %s", err)
  }

  _, err = bitly.LinkCountries(testUrl, Metrics{limit:1})
  if err != nil {
    t.Fatalf("bitly link/countries returned an error %s", err)
  }

  _, err = bitly.LinkInfo(testUrl, Metrics{limit:1})
  if err != nil {
    t.Fatalf("bitly link/info returned an error %s", err)
  }
  
  // _, err = bitly.LinkContent(testUrl, "html")
  // if err != nil {
    // t.Fatalf("bitly link/content returned an error %s", err)
  // }

  _, err = bitly.LinkCategory(testUrl)
  if err != nil {
    t.Fatalf("bitly link/category returned an error %s", err)
  }

  _, err = bitly.LinkLocation(testUrl)
  if err != nil {
    t.Fatalf("bitly link/location returned an error %s", err)
  }

  _, err = bitly.LinkSocial(testUrl)
  if err != nil {
    t.Fatalf("bitly link/social returned an error %s", err)
  }
}

func TestUserMetrics(t *testing.T) {
  bitly := getConnection(t)

  _, err := bitly.UserClicks(Metrics{limit:1})
  if err != nil {
    t.Fatalf("bitly user/clicks returned an error %s", err)
  }

  _, err = bitly.UserCountries(Metrics{limit:1})
  if err != nil {
    t.Fatalf("bitly user/countries returned an error %s", err)
  }

  _, err = bitly.UserPopularLinks(Metrics{limit:1})
  if err != nil {
    t.Fatalf("bitly user/popular_links returned an error %s", err)
  }

  _, err = bitly.UserReferrers(Metrics{limit:1})
  if err != nil {
    t.Fatalf("bitly user/referrers returned an error %s", err)
  }

  _, err = bitly.UserReferringDomains(Metrics{limit:1})
  if err != nil {
    t.Fatalf("bitly user/referrering_domains returned an error %s", err)
  }

  _, err = bitly.UserShareCounts(Metrics{limit:1})
  if err != nil {
    t.Fatalf("bitly user/share_counts returned an error %s", err)
  }

  _, err = bitly.UserShareCountsByType(Metrics{limit:1})
  if err != nil {
    t.Fatalf("bitly user/share_counts_by_share_type returned an error %s", err)
  }

  _, err = bitly.UserShortenCounts(Metrics{limit:1})
  if err != nil {
    t.Fatalf("bitly user/shorten_counts returned an error %s", err)
  }
}
