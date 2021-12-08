package time

import (
	timeUtil "time"
)

//TO THINK 面向对象的编程方式需要对属性和状态的职责梳理得清晰，要从大局思考问题
type Timer struct {
	total   int64
	isPause bool
	begin   int64
}

func NewTimer() *Timer {
	timer := &Timer{}
	timer.Init(0)
	return timer
}

func (t *Timer) Init(time int64) {
	t.total = time
	t.isPause = true
}

func (t *Timer) Start() {
	if !t.isPause {
		return
	}
	t.begin = timeUtil.Now().UnixNano() / 1e6
	t.isPause = false
}

//从指定时间开始计时
func (t *Timer) StartAt(time int64) {
	if !t.isPause {
		return
	}
	t.begin = timeUtil.Now().UnixNano() / 1e6
	t.isPause = false
	t.total = time
}

func (t *Timer) PlayMediaStartAt(time int64) {
	t.begin = timeUtil.Now().UnixNano() / 1e6
	t.isPause = false
	t.total = time
}

func (t *Timer) SeekTo(time int64) {
	t.begin = timeUtil.Now().UnixNano() / 1e6
	t.total = time
}

func (t *Timer) Pause() {
	if t.isPause {
		return
	}
	t.total += (timeUtil.Now().UnixNano()/1e6 - t.begin)
	t.isPause = true
}

func (t *Timer) PauseAt(time int64) {
	t.Init(time)
}

func (t *Timer) IsPause() bool {
	return t.isPause
}

func (t *Timer) GetBegin() int64 {
	return t.begin
}
func (t *Timer) CurrentTime() int64 {
	if t.isPause {
		return t.total
	}
	return t.total + timeUtil.Now().UnixNano()/1e6 - t.begin
}
func (t *Timer) DeltaTime() int64 {
	if t.isPause {
		return 0
	}
	return timeUtil.Now().UnixNano()/1e6 - t.begin
}


