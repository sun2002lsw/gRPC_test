# gRPC 연습하기
- 서버는 go
- 클라는 C++  

## 1. hello world
> - 그냥 hello world 문장 띄우기

## 2. echo server
> ### 2-1. unary
> - `1. hello world`와 거의 유사함. 단순한 에코 서버
> ### 2-2. server streaming
> - 클라는 그냥 전송, 서버에서 각 글자를 쪼개서 에코
> ### 2-3. client streaming
> - 클라에서 각 단어를 쪼개서 전송, 서버에서는 취합해서 에코
> ### 2-4. bidirectional streaming
> - `2-2. server streaming`와 `2-3. client streaming`를 합친것
