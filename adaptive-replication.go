type N struct{ Name string; L []int; mu sync.Mutex }

func (n *N) r(l int){ n.mu.Lock(); n.L=append(n.L,l);

 if len(n.L)>10{n.L=n.L[len(n.L) 10:]} n.mu.Unlock() }

func avg(a []int) float64{ if len(a)==0{return 1e9}; s:=0; 

for _,v:=range a{s+=v}; return float64(s)/float64(len(a)) }

func sim(name string,ch chan< struct{string;int},stop< chan 
struct{})
{
 for{
  select{ case < stop: return; default:
   l:=rand.Intn(500)+50
   ch< struct{string;int}{name,l}
   time.Sleep(time.Duration(300+rand.Intn(700))*time.Millisecond)
  }
 }
}

func pick(nodes []*N,k int)[]string{

 type p struct{n string;v float64}

 var ps []p

 for _,x:=range nodes{ x.mu.Lock(); 
ps=append(ps,p{x.Name,avg(x.L)}); x.mu.Unlock() }

 sort.Slice(ps,func(i,j int)bool{return ps[i].v<ps[j].v})

 if k>len(ps){k=len(ps)}

 out:=make([]string,0,k)

 for i:=0;i<k;i++{out=append(out,ps[i].n)}

 return out
}

func main(){

 rand.Seed(time.Now().UnixNano())

 names:=[]string{"A","B","C","D","E"}

 nodes:=make([]*N,0,len(names))

 for _,s:=range names{nodes=append(nodes,&N{Name:s})}

 ch:=make(chan struct{string;int},100)

 stop:=make(chan struct{})

 for _,s:=range names{ go sim(s,ch,stop) }

 go func(){

  t:=time.NewTicker(5*time.Second)

  for range t.C{ fmt.Println("replicas",pick(nodes,3)) }

 }()

 for ev:=range ch{

  for _,nn:=range nodes{ if nn.Name==ev.string{ nn.r(ev.int) } }

 }
}

