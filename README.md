# backlose

The command which only close the issue of backlog.

```
>.\backlose.exe 68
YOURPROJECTNAME-68 is closed.
```

Usage

```
git clone https://github.com/kongou-ae/backlose.git
cd backlose
go get github.com/kongou-ae/backlose/cmd
go build
cp .backlose.yml.sample .backlose.yml
vi .backlose.yml
./backlose [issue-num]
```