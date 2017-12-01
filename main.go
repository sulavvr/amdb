package main

import (
	// "fmt"
	"encoding/json"
	"fmt"
	"github.com/sulavvr/neogo/server"
)

func main() {
	server.Start()
	// parseCSV()
}

type test struct {
	castId    int    `json:"cast_id"`
	character string `json:"character"`
	credit    string `json:"credit_id"`
	gender    int    `json:"gender"`
	id        int    `json:"id"`
	name      string `json:"name"`
	order     int    `json:"order"`
	profile   string `json:"profile_path"`
}

func parseCSV() {
	str := `[{"cast_id": 14, "character": "Woody (voice)", "credit_id": "52fe4284c3a36847f8024f95", "gender": 2, "id": 31, "name": "Tom Hanks", "order": 0, "profile_path": "/pQFoyx7rp09CJTAb932F2g8Nlho.jpg"}, {"cast_id": 15, "character": "Buzz Lightyear (voice)", "credit_id": "52fe4284c3a36847f8024f99", "gender": 2, "id": 12898, "name": "Tim Allen", "order": 1, "profile_path": "/uX2xVf6pMmPepxnvFWyBtjexzgY.jpg"}, {"cast_id": 16, "character": "Mr. Potato Head (voice)", "credit_id": "52fe4284c3a36847f8024f9d", "gender": 2, "id": 7167, "name": "Don Rickles", "order": 2, "profile_path": "/h5BcaDMPRVLHLDzbQavec4xfSdt.jpg"}, {"cast_id": 17, "character": "Slinky Dog (voice)", "credit_id": "52fe4284c3a36847f8024fa1", "gender": 2, "id": 12899, "name": "Jim Varney", "order": 3, "profile_path": "/eIo2jVVXYgjDtaHoF19Ll9vtW7h.jpg"}, {"cast_id": 18, "character": "Rex (voice)", "credit_id": "52fe4284c3a36847f8024fa5", "gender": 2, "id": 12900, "name": "Wallace Shawn", "order": 4, "profile_path": "/oGE6JqPP2xH4tNORKNqxbNPYi7u.jpg"}, {"cast_id": 19, "character": "Hamm (voice)", "credit_id": "52fe4284c3a36847f8024fa9", "gender": 2, "id": 7907, "name": "John Ratzenberger", "order": 5, "profile_path": "/yGechiKWL6TJDfVE2KPSJYqdMsY.jpg"}, {"cast_id": 20, "character": "Bo Peep (voice)", "credit_id": "52fe4284c3a36847f8024fad", "gender": 1, "id": 8873, "name": "Annie Potts", "order": 6, "profile_path": "/eryXT84RL41jHSJcMy4kS3u9y6w.jpg"}, {"cast_id": 26, "character": "Andy (voice)", "credit_id": "52fe4284c3a36847f8024fc1", "gender": 0, "id": 1116442, "name": "John Morris", "order": 7, "profile_path": "/vYGyvK4LzeaUCoNSHtsuqJUY15M.jpg"}, {"cast_id": 22, "character": "Sid (voice)", "credit_id": "52fe4284c3a36847f8024fb1", "gender": 2, "id": 12901, "name": "Erik von Detten", "order": 8, "profile_path": "/twnF1ZaJ1FUNUuo6xLXwcxjayBE.jpg"}, {"cast_id": 23, "character": "Mrs. Davis (voice)", "credit_id": "52fe4284c3a36847f8024fb5", "gender": 1, "id": 12133, "name": "Laurie Metcalf", "order": 9, "profile_path": "/unMMIT60eoBM2sN2nyR7EZ2BvvD.jpg"}, {"cast_id": 24, "character": "Sergeant (voice)", "credit_id": "52fe4284c3a36847f8024fb9", "gender": 2, "id": 8655, "name": "R. Lee Ermey", "order": 10, "profile_path": "/r8GBqFBjypLUP9VVqDqfZ7wYbSs.jpg"}, {"cast_id": 25, "character": "Hannah (voice)", "credit_id": "52fe4284c3a36847f8024fbd", "gender": 1, "id": 12903, "name": "Sarah Freeman", "order": 11, "profile_path": None}, {"cast_id": 27, "character": "TV Announcer (voice)", "credit_id": "52fe4284c3a36847f8024fc5", "gender": 2, "id": 37221, "name": "Penn Jillette", "order": 12, "profile_path": "/zmAaXUdx12NRsssgHbk1T31j2x9.jpg"}]`

	fmt.Println(json.Valid([]byte(str)))
	// x := test{}
	// _ = json.Unmarshal([]byte(str), &x)
	// fmt.Println(x)
}
