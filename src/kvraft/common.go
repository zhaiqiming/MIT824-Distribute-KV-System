package raftkv

const (
	OK             = "OK"
	ErrNoKey       = "ErrNoKey"
	ErrTimeout     = "ErrTimeout"
	ErrWrongLeader = "ErrWrongLeader"
)

type Err string

// Put or Append
type PutAppendArgs struct {
	Key   string
	Value string
	Op    string // "Put" or "Append"
	// You'll have to add definitions here.
	// Field names must start with capital letters,
	// otherwise RPC will break.
	ClientId int64
	Seq      int
}

type PutAppendReply struct {
	WrongLeader bool
	Err         Err
}

type GetArgs struct {
	Key string
	// You'll have to add definitions here.
	ClientId int64
	Seq      int
}

type GetReply struct {
	WrongLeader bool
	Err         Err
	Value       string
}