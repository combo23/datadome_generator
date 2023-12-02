package datadome

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/combo23/datadome_generator/internal/timezone"
)

func (t *DataDomePayload) GetJSData() (string, error) {
	var jsData jsDataCH
	if t.JsType == "ch" {
		jsData = jsDataCH{}
	} else if t.JsType == "le" {
		jsData = t.jsdata
	}
	if t.JsType == "ch" {
		jsData.Ttst = math.Floor(t.rand.Float64()*80) + t.rand.Float64()                                                                                                                                                          //some performance.now(), doesnt matter
		jsData.Tagpu = math.Floor(t.rand.Float64()*15) + t.rand.Float64()                                                                                                                                                         //some performance.now(), doesnt matter
		jsData.Jset = int(time.Now().Unix())                                                                                                                                                                                      //just time.now()
		jsData.Ua = t.UserAgent                                                                                                                                                                                                   //user agent
		jsData.Ifov = false                                                                                                                                                                                                       //always false
		screendata := resolutions[t.rand.Intn(len(resolutions))]                                                                                                                                                                  //generation of screendata
		jsData.RsH = screendata.Height                                                                                                                                                                                            //height
		jsData.RsW = screendata.Width                                                                                                                                                                                             //width
		jsData.ArsH = t.rand.Intn(screendata.Height*3) / 2                                                                                                                                                                        //availHeight
		jsData.ArsW = t.rand.Intn(screendata.Width*3) / 2                                                                                                                                                                         //availWidth
		jsData.BrOh = t.rand.Intn(screendata.Height*3) / 2                                                                                                                                                                        //outerHeight
		jsData.BrOw = t.rand.Intn(screendata.Width*3) / 2                                                                                                                                                                         //outerWidth
		jsData.BrH = t.rand.Intn(screendata.Height*3) / 2                                                                                                                                                                         //innerHeight
		jsData.BrW = t.rand.Intn(screendata.Width*3) / 2                                                                                                                                                                          //innerWidth
		jsData.Phe = false                                                                                                                                                                                                        //always false
		jsData.Nm = false                                                                                                                                                                                                         //always false
		jsData.StrSs = true                                                                                                                                                                                                       //always true
		jsData.StrLs = true                                                                                                                                                                                                       //always true
		jsData.StrIdb = true                                                                                                                                                                                                      //always true
		jsData.StrOdb = false                                                                                                                                                                                                     //always true
		jsData.Mmt = "application/pdf,text/pdf"                                                                                                                                                                                   //always this
		jsData.Plg = 5                                                                                                                                                                                                            //number of plugins
		jsData.Plu = "PDF Viewer,Chrome PDF Viewer,Chromium PDF Viewer,Microsoft Edge PDF Viewer,WebKit built-in PDF"                                                                                                             //idk why but it always shows only this plugins so its static ig
		jsData.Lg = "en-GB"                                                                                                                                                                                                       //language
		jsData.Glvd = possibleGlvd[t.rand.Intn(len(possibleGlvd))]                                                                                                                                                                //glvd
		jsData.Glrd = possibleGlrd[jsData.Glvd][t.rand.Intn(len(possibleGlrd))]                                                                                                                                                   //glvd !!!!NEED TO GATHER MORE DATA!!!!
		jsData.Cvs = true                                                                                                                                                                                                         //always true
		jsData.Opts = "ajaxListenerPath"                                                                                                                                                                                          //99% thats static xd
		jsData.RsCd = possibleRs_cd[t.rand.Intn(len(possibleRs_cd))]                                                                                                                                                              //screen.colorDepth
		jsData.Vnd = "Google Inc."                                                                                                                                                                                                //navigator.vendor static
		jsData.Dvm = possibleDvm[t.rand.Intn(len(possibleDvm))]                                                                                                                                                                   //deviceMemory
		jsData.Hc = possibleHc[t.rand.Intn(len(possibleHc))]                                                                                                                                                                      //hardwareConcurrency
		jsData.Pr = possiblePixelRatios[t.rand.Intn(len(possiblePixelRatios))]                                                                                                                                                    //devicePixelRatio
		jsData.Dp0 = false                                                                                                                                                                                                        //also static
		jsData.Nddc = rand.Intn(3)                                                                                                                                                                                                //0, 1 ,2
		jsData.Jsf = false                                                                                                                                                                                                        //static
		jsData.Wdifrm = false                                                                                                                                                                                                     //always false
		jsData.Npmtm = randomBoolean(t.rand)                                                                                                                                                                                      //random boolean value
		jsData.Ftsovdr = false                                                                                                                                                                                                    //static
		jsData.Ftsovdr2 = false                                                                                                                                                                                                   //static
		jsData.Lb = false                                                                                                                                                                                                         //static
		jsData.Wbd = false                                                                                                                                                                                                        //is chromedriver xddd
		jsData.Plgod = false                                                                                                                                                                                                      //static
		jsData.Plgne = true                                                                                                                                                                                                       //static
		jsData.Plgre = true                                                                                                                                                                                                       //basically the same as plgne
		jsData.Plgof = false                                                                                                                                                                                                      //static also
		jsData.Plggt = false                                                                                                                                                                                                      //static
		jsData.Pltod = false                                                                                                                                                                                                      //static
		jsData.Hcovdr = false                                                                                                                                                                                                     //static
		jsData.Hcovdr2 = false                                                                                                                                                                                                    //static
		jsData.Plovdr = false                                                                                                                                                                                                     //static
		jsData.Plovdr2 = false                                                                                                                                                                                                    //static
		jsData.Lo = false                                                                                                                                                                                                         //static
		jsData.TsMtp = 0                                                                                                                                                                                                          //static
		jsData.TsTec = false                                                                                                                                                                                                      //static
		jsData.TsTsa = false                                                                                                                                                                                                      //static
		jsData.Bid = "NA"                                                                                                                                                                                                         //static
		jsData.Hdn = false                                                                                                                                                                                                        //!!document.hidden
		jsData.Awe = false                                                                                                                                                                                                        //!!window.awesomium
		jsData.Geb = false                                                                                                                                                                                                        //!!window.geb (some custom shit, normal browsers dont have this)
		jsData.Dat = false                                                                                                                                                                                                        //'domAutomation' in window || 'domAutomationController' in window
		jsData.Med = "defined"                                                                                                                                                                                                    //window.navigator.mediaDevices
		jsData.Aco = "probably"                                                                                                                                                                                                   //canPlayType('audio/ogg; codecs="vorbis"')
		jsData.Acots = false                                                                                                                                                                                                      //isTypeSupported('audio/ogg; codecs="vorbis"')
		jsData.Acmp = "probably"                                                                                                                                                                                                  //canPlayType('audio/mpeg;')
		jsData.Acmpts = true                                                                                                                                                                                                      //isTypeSupported('audio/mpeg;"')
		jsData.Acw = "probably"                                                                                                                                                                                                   //canPlayType('audio/wav; codecs="1"')
		jsData.Acwts = false                                                                                                                                                                                                      //isTypeSupported('audio/wav; codecs="1"') !!CHECK THIS SEEMS WEIRD ASF
		jsData.Cokys = "bG9hZFRpbWVzY3NpYXBwL="                                                                                                                                                                                   //base64 static string
		jsData.Sqt = false                                                                                                                                                                                                        //window.external && window.external.toString && window.external.toString().indexOf('Sequentum') > -1
		jsData.Cfpfe = "RXJyb3I6IENhbm5vdCByZWFkIHByb3BlcnRpZXMgb2YgbnVsbCAocmVhZGluZyAndG9TdHJpbmcnKQ=="                                                                                                                         //static stack error
		jsData.Stcfp = "cHR0ZW1wbGF0ZXMvMjAyMzEwLjEuMC9vdEJhbm5lclNkay5qczo3OjQxODkxNikKICAgIGF0IEhUTUxIdG1sRWxlbWVudC5yIChodHRwczovL2Nkbi5jb29raWVsYXcub3JnL3NjcmlwdHRlbXBsYXRlcy8yMDIzMTAuMS4wL290QmFubmVyU2RrLmpzOjc6Njk3MTkp" //static stack error
		jsData.Usb = "defined"                                                                                                                                                                                                    //self explanatory
		jsData.Acma = "maybe"                                                                                                                                                                                                     //canPlayType('audio/x-m4a;')
		jsData.Acmats = false                                                                                                                                                                                                     //isTypeSupported('audio/x-m4a;')
		jsData.Acaa = "probably"                                                                                                                                                                                                  //canPlayType('audio/aac;')
		jsData.Acaats = true                                                                                                                                                                                                      //isTypeSupported('audio/aac;')
		jsData.Ac3 = ""                                                                                                                                                                                                           //canPlayType('audio/3gpp;')
		jsData.Ac3Ts = false                                                                                                                                                                                                      //isTypeSupported('audio/3gpp;')
		jsData.Acf = "probably"                                                                                                                                                                                                   //canPlayType('audio/flac;')
		jsData.Acfts = false                                                                                                                                                                                                      //isTypeSupported('audio/flac;')
		jsData.Acmp4 = "maybe"                                                                                                                                                                                                    //canPlayType('audio/mp4;')
		jsData.Acmp4Ts = false                                                                                                                                                                                                    //isTypeSupported('audio/mp4;')
		jsData.Acmp3 = "maybe"                                                                                                                                                                                                    //canPlayType('audio/mp3;')
		jsData.Acmp3Ts = false                                                                                                                                                                                                    //isTypeSupported('audio/mp3;')
		jsData.Acwm = "maybe"                                                                                                                                                                                                     //canPlayType('audio/webm;')
		jsData.Acwmts = false                                                                                                                                                                                                     //isTypeSupported('audio/webm;')
		jsData.Ocpt = false                                                                                                                                                                                                       //-1 === _0x521bea.canPlayType.toString().indexOf('canPlayType')
		jsData.Vco = "probably"                                                                                                                                                                                                   //canPlayType('video/ogg; codecs="theora"')
		jsData.Vcots = false                                                                                                                                                                                                      //isTypeSupported('video/ogg; codecs="theora"')
		jsData.Vch = "probably"                                                                                                                                                                                                   //canPlayType('video/mp4; codecs="avc1.42E01E"')
		jsData.Vchts = true                                                                                                                                                                                                       //isTypeSupported('video/mp4; codecs="avc1.42E01E"')
		jsData.Vcw = "probably"                                                                                                                                                                                                   //canPlayType('video/webm; codecs="vp8, vorbis"')
		jsData.Vcwts = true                                                                                                                                                                                                       //isTypeSupported('video/webm; codecs="vp8, vorbis"')
		jsData.Vc3 = "maybe"                                                                                                                                                                                                      //canPlayType('video/3gpp;')
		jsData.Vc3Ts = false                                                                                                                                                                                                      //isTypeSupported('video/3gpp;')
		jsData.Vcmp = ""                                                                                                                                                                                                          //canPlayType('video/mpeg;')
		jsData.Vcmpts = false                                                                                                                                                                                                     //isTypeSupported('video/mpeg')
		jsData.Vcq = ""                                                                                                                                                                                                           //canPlayType('video/quicktime;')
		jsData.Vcqts = false                                                                                                                                                                                                      //isTypeSupported('video/quicktime;')
		jsData.Vc1 = "probably"                                                                                                                                                                                                   //canPlayType('video/mp4; codecs="av01.0.08M.08"')
		jsData.Vc1Ts = true                                                                                                                                                                                                       //isTypeSupported('video/mp4; codecs="av01.0.08M.08"')
		jsData.So = possibleOrientationTypes[t.rand.Intn(len(possibleOrientationTypes))]                                                                                                                                          //orientation type
		jsData.Psn = true                                                                                                                                                                                                         //!!window.PermissionStatus && window.PermissionStatus.prototype.hasOwnProperty('name')
		jsData.Edp = true                                                                                                                                                                                                         //!!window.EyeDropper
		jsData.Addt = true                                                                                                                                                                                                        //!!window.AudioData
		jsData.Wsdc = true                                                                                                                                                                                                        //!!window.WritableStreamDefaultController
		jsData.Ccsr = true                                                                                                                                                                                                        //!!window.CSSCounterStyleRule
		jsData.Nuad = true                                                                                                                                                                                                        //!!window.NavigatorUAData
		jsData.Bcda = false                                                                                                                                                                                                       //!!window.BarcodeDetector (what fucking computer has barcode scanner)
		jsData.Bfr = false                                                                                                                                                                                                        //!!window.Buffer
		jsData.Dbov = false                                                                                                                                                                                                       //!!('' + window.console.debug).match(/[\)\( ]{3}[>= ]{3}\{\n[ r]{9}etu[n r]{3}n[lu]{3}/ ))
		jsData.Ecpc = false                                                                                                                                                                                                       //some long if statement checking for electron substring
		jsData.Lgs = randomBoolean(t.rand)                                                                                                                                                                                        //_0x4908b3.lgs = '' !== navigator.languages (very weird comparison between string and a list so ig its random)
		jsData.Lgsod = randomBoolean(t.rand)                                                                                                                                                                                      //!!Object.getOwnPropertyDescriptor(navigator,'languages')
		jsData.Idn = true                                                                                                                                                                                                         //!(!window.Intl || !Intl.DisplayNames)
		jsData.Capi = false                                                                                                                                                                                                       //!!(window.navigator && window.navigator.contacts && window.navigator.ContactsManager)
		jsData.Svde = false                                                                                                                                                                                                       //!!window.SVGDiscardElement
		jsData.Vpbq = true                                                                                                                                                                                                        //!!(window.HTMLVideoElement &&window.HTMLVideoElement.prototype &&window.HTMLVideoElement.prototype.getVideoPlaybackQuality)
		jsData.Ucdv = false                                                                                                                                                                                                       //long random if statement
		jsData.Spwn = false                                                                                                                                                                                                       //!!window.spawn
		jsData.Emt = false                                                                                                                                                                                                        //!!window.emit
		jsData.Wdif = false                                                                                                                                                                                                       //!!_0x57c363.contentWindow.navigator.webdriver
		jsData.Wdw = true                                                                                                                                                                                                         //static
		jsData.Prm = true                                                                                                                                                                                                         //also static long if statemnet
		jsData.Wwl = false                                                                                                                                                                                                        //for now its static, need to be changed

		eva, err := getEvavalue(t.UserAgent)
		if err != nil {
			return "", err
		}
		jsData.Eva = eva

		if testing.Verbose() {
			//for testing purposes
			jsData.Tzp = "Europe/London"
			jsData.Tz = 0
		} else {
			tzp, err := timezone.GetTimezone(t.IP)
			if err != nil {
				return "", err
			}
			jsData.Tzp = tzp.TimeZone.ID
			jsData.Tz = -(tzp.TimeZone.UTCOffset / 60)
		}

	} else if t.JsType == "le" {
		jsData.jsDataLe = new(jsDataLe)
		jsData.Jset = jsData.Jset + rand.Intn(20) //overriding time.now() from ch
		jsData.Dcok = fmt.Sprintf(".%v", strings.ToLower(t.Host))
		jsData.MMC = t.events.Mousemove
		jsData.MSC = t.events.Scroll
		jsData.MCC = t.events.Click
		jsData.MCmR = calculateMcmr(t.events.Mousemove, t.events.Click)
		jsData.MMsR = calculateMmsr(t.events.Mousemove, t.events.Scroll)
		data, err := t.generateMousemovements()
		if err != nil {
			return "", err
		}
		jsData.EsSigmdn = data.EsSigmdn
		jsData.EsMumdn = data.EsMumdn
		jsData.EsDistmdn = data.EsDistmdn
		jsData.EsAngsmdn = data.EsAngsmdn
		jsData.EsAngemdn = data.EsAngemdn
	}

	data, err := json.Marshal(jsData)
	if err != nil {
		return "", err
	}
	//TODO
	t.jsdata = jsData
	return string(data), nil
}

type jsDataCH struct {
	Opts     string  `json:"opts"`
	Ttst     float64 `json:"ttst"`
	Ifov     bool    `json:"ifov"`
	Dp0      bool    `json:"dp0"`
	Tagpu    float64 `json:"tagpu"`
	Glvd     string  `json:"glvd"`
	Glrd     string  `json:"glrd"`
	Hc       int     `json:"hc"`
	BrOh     int     `json:"br_oh"`
	BrOw     int     `json:"br_ow"`
	Ua       string  `json:"ua"`
	Wbd      bool    `json:"wbd"`
	Wdif     bool    `json:"wdif"`
	Wdifrm   bool    `json:"wdifrm"`
	Npmtm    bool    `json:"npmtm"`
	BrH      int     `json:"br_h"`
	BrW      int     `json:"br_w"`
	Nddc     int     `json:"nddc"`
	RsH      int     `json:"rs_h"`
	RsW      int     `json:"rs_w"`
	RsCd     int     `json:"rs_cd"`
	Phe      bool    `json:"phe"`
	Nm       bool    `json:"nm"`
	Jsf      bool    `json:"jsf"`
	Lg       string  `json:"lg"`
	Pr       int     `json:"pr"`
	ArsH     int     `json:"ars_h"`
	ArsW     int     `json:"ars_w"`
	Tz       int     `json:"tz"`
	StrSs    bool    `json:"str_ss"`
	StrLs    bool    `json:"str_ls"`
	StrIdb   bool    `json:"str_idb"`
	StrOdb   bool    `json:"str_odb"`
	Plgod    bool    `json:"plgod"`
	Plg      int     `json:"plg"`
	Plgne    bool    `json:"plgne"`
	Plgre    bool    `json:"plgre"`
	Plgof    bool    `json:"plgof"`
	Plggt    bool    `json:"plggt"`
	Pltod    bool    `json:"pltod"`
	Hcovdr   bool    `json:"hcovdr"`
	Hcovdr2  bool    `json:"hcovdr2"`
	Plovdr   bool    `json:"plovdr"`
	Plovdr2  bool    `json:"plovdr2"`
	Ftsovdr  bool    `json:"ftsovdr"`
	Ftsovdr2 bool    `json:"ftsovdr2"`
	Lb       bool    `json:"lb"`
	Eva      int     `json:"eva"`
	Lo       bool    `json:"lo"`
	TsMtp    int     `json:"ts_mtp"`
	TsTec    bool    `json:"ts_tec"`
	TsTsa    bool    `json:"ts_tsa"`
	Vnd      string  `json:"vnd"`
	Bid      string  `json:"bid"`
	Mmt      string  `json:"mmt"`
	Plu      string  `json:"plu"`
	Hdn      bool    `json:"hdn"`
	Awe      bool    `json:"awe"`
	Geb      bool    `json:"geb"`
	Dat      bool    `json:"dat"`
	Med      string  `json:"med"`
	Aco      string  `json:"aco"`
	Acots    bool    `json:"acots"`
	Acmp     string  `json:"acmp"`
	Acmpts   bool    `json:"acmpts"`
	Acw      string  `json:"acw"`
	Acwts    bool    `json:"acwts"`
	Acma     string  `json:"acma"`
	Acmats   bool    `json:"acmats"`
	Acaa     string  `json:"acaa"`
	Acaats   bool    `json:"acaats"`
	Ac3      string  `json:"ac3"`
	Ac3Ts    bool    `json:"ac3ts"`
	Acf      string  `json:"acf"`
	Acfts    bool    `json:"acfts"`
	Acmp4    string  `json:"acmp4"`
	Acmp4Ts  bool    `json:"acmp4ts"`
	Acmp3    string  `json:"acmp3"`
	Acmp3Ts  bool    `json:"acmp3ts"`
	Acwm     string  `json:"acwm"`
	Acwmts   bool    `json:"acwmts"`
	Ocpt     bool    `json:"ocpt"`
	Vco      string  `json:"vco"`
	Vcots    bool    `json:"vcots"`
	Vch      string  `json:"vch"`
	Vchts    bool    `json:"vchts"`
	Vcw      string  `json:"vcw"`
	Vcwts    bool    `json:"vcwts"`
	Vc3      string  `json:"vc3"`
	Vc3Ts    bool    `json:"vc3ts"`
	Vcmp     string  `json:"vcmp"`
	Vcmpts   bool    `json:"vcmpts"`
	Vcq      string  `json:"vcq"`
	Vcqts    bool    `json:"vcqts"`
	Vc1      string  `json:"vc1"`
	Vc1Ts    bool    `json:"vc1ts"`
	Dvm      int     `json:"dvm"`
	Sqt      bool    `json:"sqt"`
	So       string  `json:"so"`
	Wdw      bool    `json:"wdw"`
	Cokys    string  `json:"cokys"`
	Ecpc     bool    `json:"ecpc"`
	Lgs      bool    `json:"lgs"`
	Lgsod    bool    `json:"lgsod"`
	Psn      bool    `json:"psn"`
	Edp      bool    `json:"edp"`
	Addt     bool    `json:"addt"`
	Wsdc     bool    `json:"wsdc"`
	Ccsr     bool    `json:"ccsr"`
	Nuad     bool    `json:"nuad"`
	Bcda     bool    `json:"bcda"`
	Idn      bool    `json:"idn"`
	Capi     bool    `json:"capi"`
	Svde     bool    `json:"svde"`
	Vpbq     bool    `json:"vpbq"`
	Ucdv     bool    `json:"ucdv"`
	Spwn     bool    `json:"spwn"`
	Emt      bool    `json:"emt"`
	Bfr      bool    `json:"bfr"`
	Dbov     bool    `json:"dbov"`
	Cfpfe    string  `json:"cfpfe"`
	Stcfp    string  `json:"stcfp"`
	Prm      bool    `json:"prm"`
	Tzp      string  `json:"tzp"`
	Cvs      bool    `json:"cvs"`
	Usb      string  `json:"usb"`
	Wwl      bool    `json:"wwl"`
	Jset     int     `json:"jset"`
	*jsDataLe
}

type jsDataLe struct {
	Dcok      string  `json:"dcok"`
	EsSigmdn  float64 `json:"es_sigmdn"`
	EsMumdn   float64 `json:"es_mumdn"`
	EsDistmdn float64 `json:"es_distmdn"`
	EsAngsmdn float64 `json:"es_angsmdn"`
	EsAngemdn float64 `json:"es_angemdn"`
	MSC       int     `json:"m_s_c"`
	MMC       int     `json:"m_m_c"`
	MCC       int     `json:"m_c_c"`
	MCmR      int     `json:"m_cm_r"`
	MMsR      int     `json:"m_ms_r"`
}
