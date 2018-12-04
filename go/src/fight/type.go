package fight

import "sync"

type (
	Point struct {
		X int `json:"x"`
		Y int `json:"y"`
	}

	fPoint struct {
		x, y, m1 int
		m2       byte
	}

	fUser struct {
		id            byte // 在room中的id
		score, energy int  // score => cell的个数
		userName      string
		userToken     string
		roomToken     string
		status        int // wait playing lose win
		baseLoc       Point
	}

	fQueryCounter struct {
		mu    sync.Mutex
		count int
	}

	fOpts struct {
		mu         sync.RWMutex
		accId      byte // 当前最小未分配ID (严格递增)
		playerNum  int
		row, col   int
		barbackNum int
		portalNum  int
		barrierNum int
		playing    bool
		roomToken  string
		userInfo   map[byte]*fUser // [id]*fUser
		m          *fMap
	}

	// m2 => 8字节
	// 前三位地形类型 后五位阵营
	fMap struct {
		mu        sync.RWMutex
		roomToken string
		m1        [default_max_row][default_max_col]int
		m2        [default_max_row][default_max_col]byte

		base, barback, portal, barrier []Point
	}

	/* net */
	JoinRet struct {
		Id        byte
		PlayerNum int
		Row, Col  int
		RoomToken string
	}

	EyeShot struct {
		M1 [default_eye_level*2 + 1][default_eye_level*2 + 1]int  `json:"m1"`
		M2 [default_eye_level*2 + 1][default_eye_level*2 + 1]byte `json:"m2"`
	}
	/* net */
)
