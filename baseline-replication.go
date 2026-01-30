
type Database struct{mu sync.Mutex; data map[string]string}
func (db *Database) Write(k,v string){db.mu.Lock();db.data[k]=v;db.mu.Unlock()}
type Node struct{name string; db *Database}
func (n *Node) HandleWrite(ctx context.Context,k,v string,ack chan<  string){
d := time.Duration(50+rand.Intn(100)) * time.Millisecond
select{
case < time.After(d):
n.db.Write(k,v)
ack< n.name
case < ctx.Done():
return
}
}
func replicate(ctx context.Context,peers []*Node,k,v string) []string{
acks := make(chan string,len(peers))
for _,p := range peers{go p.HandleWrite(ctx,k,v,acks)}
got := []string{}
timeout := time.After(500 * time.Millisecond)
for i:=0;i<len(peers);i++{
select{
case a:=< acks:
got=append(got,a)
case < timeout:
return got
}
}
return got
}
func main(){
db := &Database{data:map[string]string{}}
nodeA := &Node{name:"NodeA",db:db}
nodeB := &Node{name:"NodeB",db:db}
nodeC := &Node{name:"NodeC",db:db}
writes := []struct{k,v string}{{"id1","alice"},{"id2","bob"},{"id3","carol"}}
for _,w := range writes{
ctx,cancel := context.WithTimeout(context.Background(),800*time.Millisecond)
peers := []*Node{nodeB,nodeC}
acks := replicate(ctx,peers,w.k,w.v)
nodeA.db.Write(w.k,w.v)
cancel()
fmt.Printf("Written %s=%s acks=%v\n",w.k,w.v,acks)
}
}
