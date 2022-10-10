package requester

import (
	"encoding/json"
	"fmt"

	"io/ioutil"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
	//"github.com/tidwall/gjson"
)

type StockPicker struct {
	MetaData struct {
		OneInformation     string `json:"1. Information"`
		TwoSymbol          string `json:"2. Symbol"`
		ThreeLastRefreshed string `json:"3. Last Refreshed"`
		FourOutputSize     string `json:"4. Output Size"`
		FiveTimeZone       string `json:"5. Time Zone"`
	} `json:"Meta Data"`
	TimeSeriesDaily struct {
		Two0221007 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-10-07"`
		Two0221006 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-10-06"`
		Two0221005 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-10-05"`
		Two0221004 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-10-04"`
		Two0221003 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-10-03"`
		Two0220930 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-09-30"`
		Two0220929 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-09-29"`
		Two0220928 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-09-28"`
		Two0220927 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-09-27"`
		Two0220926 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-09-26"`
		Two0220923 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-09-23"`
		Two0220922 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-09-22"`
		Two0220921 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-09-21"`
		Two0220920 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-09-20"`
		Two0220919 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-09-19"`
		Two0220916 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-09-16"`
		Two0220915 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-09-15"`
		Two0220914 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-09-14"`
		Two0220913 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-09-13"`
		Two0220912 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-09-12"`
		Two0220909 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-09-09"`
		Two0220908 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-09-08"`
		Two0220907 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-09-07"`
		Two0220906 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-09-06"`
		Two0220902 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-09-02"`
		Two0220901 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-09-01"`
		Two0220831 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-08-31"`
		Two0220830 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-08-30"`
		Two0220829 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-08-29"`
		Two0220826 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-08-26"`
		Two0220825 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-08-25"`
		Two0220824 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-08-24"`
		Two0220823 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-08-23"`
		Two0220822 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-08-22"`
		Two0220819 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-08-19"`
		Two0220818 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-08-18"`
		Two0220817 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-08-17"`
		Two0220816 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-08-16"`
		Two0220815 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-08-15"`
		Two0220812 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-08-12"`
		Two0220811 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-08-11"`
		Two0220810 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-08-10"`
		Two0220809 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-08-09"`
		Two0220808 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-08-08"`
		Two0220805 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-08-05"`
		Two0220804 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-08-04"`
		Two0220803 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-08-03"`
		Two0220802 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-08-02"`
		Two0220801 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-08-01"`
		Two0220729 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-07-29"`
		Two0220728 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-07-28"`
		Two0220727 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-07-27"`
		Two0220726 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-07-26"`
		Two0220725 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-07-25"`
		Two0220722 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-07-22"`
		Two0220721 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-07-21"`
		Two0220720 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-07-20"`
		Two0220719 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-07-19"`
		Two0220718 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-07-18"`
		Two0220715 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-07-15"`
		Two0220714 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-07-14"`
		Two0220713 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-07-13"`
		Two0220712 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-07-12"`
		Two0220711 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-07-11"`
		Two0220708 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-07-08"`
		Two0220707 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-07-07"`
		Two0220706 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-07-06"`
		Two0220705 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-07-05"`
		Two0220701 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-07-01"`
		Two0220630 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-06-30"`
		Two0220629 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-06-29"`
		Two0220628 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-06-28"`
		Two0220627 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-06-27"`
		Two0220624 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-06-24"`
		Two0220623 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-06-23"`
		Two0220622 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-06-22"`
		Two0220621 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-06-21"`
		Two0220617 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-06-17"`
		Two0220616 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-06-16"`
		Two0220615 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-06-15"`
		Two0220614 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-06-14"`
		Two0220613 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-06-13"`
		Two0220610 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-06-10"`
		Two0220609 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-06-09"`
		Two0220608 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-06-08"`
		Two0220607 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-06-07"`
		Two0220606 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-06-06"`
		Two0220603 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-06-03"`
		Two0220602 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-06-02"`
		Two0220601 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-06-01"`
		Two0220531 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-05-31"`
		Two0220527 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-05-27"`
		Two0220526 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-05-26"`
		Two0220525 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-05-25"`
		Two0220524 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-05-24"`
		Two0220523 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-05-23"`
		Two0220520 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-05-20"`
		Two0220519 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-05-19"`
		Two0220518 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-05-18"`
		Two0220517 struct {
			OneOpen    string `json:"1. open"`
			TwoHigh    string `json:"2. high"`
			ThreeLow   string `json:"3. low"`
			FourClose  string `json:"4. close"`
			FiveVolume string `json:"5. volume"`
		} `json:"2022-05-17"`
	} `json:"Time Series (Daily)"`
}

func MakeRequest(symbl string) (data StockPicker) {
	log.SetLevel(log.DebugLevel)
	req, _ := http.NewRequest("GET", "https://www.alphavantage.co/query?", nil)

	q := req.URL.Query()
	q.Add("apikey", "C227WD9W3LUVKVV9")
	q.Add("function", "TIME_SERIES_DAILY")
	q.Add("symbol", "MSFT")
	req.URL.RawQuery = q.Encode()

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("error occurrred at get")
	}

	defer resp.Body.Close()

	// fmt.Println(req.URL.String())

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error occurred")
	}

	//var data StockPicker

	err = json.Unmarshal(body, &data)

	if err != nil {
		panic(err)
	}

	return data

}
