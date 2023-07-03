package registry

type EventType int

const Add EventType = 0
const Update EventType = 1
const Delete EventType = 2
const Select EventType = 3

type AsyncResultCode int

const Success AsyncResultCode = 0
const Error AsyncResultCode = 1
const TimeOut AsyncResultCode = 2
const UnKnowError AsyncResultCode = 999
