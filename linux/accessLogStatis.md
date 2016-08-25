
找出502的记录
> tail -3333 /home/wwwlogs/access.log | awk 'BEGIN{FS=" "}; {if(504== $9){print $0}}'

linux统计某个文件内容的 每个值出现的次
> last | cut -d ' ' -f 1 | sort | uniq -c | sort -t ' ' -k 1 -n -r


