# myhead
Command line to output two files
## command
myhead -s=1 -e=3 -d=true hoge.txt huga.txt

## option
 - -s int = start point
 - -e int = end point
 - -d bool = output two files


```
$ go run myhead.go -s=0 -e=3 -d=false hoge.txt huga.txt

<<<<<<hoge.txt>>>>>>>
11月23日　 半年振りに書く。今日は祝日でした。普通なら嬉しいはずの休みが嬉しくないのは、

　　　　　　　　　同じ職場の片思いしてる女性に会えないから。いい年して餓鬼みたいな純粋さを見せてもキモいだけだが。
<<<<<<huga.txt>>>>>>>
Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod 
tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, 
quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. 
sakaguchirionanoMacBook-puro:myhead sakaguchiriona$ 
```
