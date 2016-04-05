package main
import (
  "flag"
  "os"
  "fmt"
  "github.com/parnurzeal/gorequest"
  "encoding/json"
  "appstore-hipchat-integration/models"
  "github.com/tbruyelle/hipchat-go/hipchat"
)

var (
  token  = flag.String("token", "", "The HipChat AuthToken")
  roomId = flag.String("room", "", "The HipChat room id")
  test   = flag.Bool("t", false, "Enable auth_test parameter")
)

func main() {
  flag.Parse()
  if *token == "" || *roomId == "" {
    flag.PrintDefaults()
    return
  }
  hipchat.AuthTest = *test
  c := hipchat.NewClient(*token)

  url := "https://itunes.apple.com/au/rss/customerreviews/id=1029783189/page=1/json"
  resp, body, httperr := gorequest.New().Get(url).End()
  if httperr == nil {
    fmt.Println(resp)
    feed := models.AutoGenerated{}
    json.Unmarshal([]byte(body), &feed)
    notifRq := &hipchat.NotificationRequest{Message: "<h1>" + feed.Feed.Entry[1].Title.Label + "</h1> (" + feed.Feed.Entry[1].ImReleaseDate.Label + ")</br> " + feed.Feed.Entry[1].Content.Label}
    //notifRq := &hipchat.NotificationRequest{Message: "<table border=\"0\" width=\"100%\"> <tr> <td> <table border=\"0\" width=\"100%\" cellspacing=\"0\" cellpadding=\"0\"> <tr valign=\"top\" align=\"left\"> <td width=\"100%\"> <b><a href=\"https://itunes.apple.com/au/app/william-hill-home-betting/id1029783189?mt=8&uo=2\">Betslip needs upgrade</a></b><br/> <font size=\"2\" face=\"Helvetica,Arial,Geneva,Swiss,SunSans-Regular\"> </font> </td> </tr> </table> </td> </tr> <tr> <td> <font size=\"2\" face=\"Helvetica,Arial,Geneva,Swiss,SunSans-Regular\"><br/>I have to get onto the website to put the type of bets I like on. No lucky 63&#39;s etc types available through the app.</font><br/> </td> </tr> </table>"}
    resp, err := c.Room.Notification(*roomId, notifRq)
    if err != nil {
      fmt.Fprintf(os.Stderr, "Error during room notification %q\n", err)
      fmt.Fprintf(os.Stderr, "Server returns %+v\n", resp)
      return
    }

    if hipchat.AuthTest {
      _, ok := hipchat.AuthTestResponse["success"]
      fmt.Println("Authentification succeed :", ok)
    } else {
      fmt.Println("Lol sent !")
    }
  }
}