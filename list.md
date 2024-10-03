- **auths**
  - `auths-api`=8888
  - `auths-service`=10000
  - `auths-service-health`=20000
  - `auths-api-health`=30000

- **user**
  - `user-api`=8889
  - `user-service`=10001
  - `user-service-health`=20001
  - `user-api-health`=30001
  - `depend`=[auths,]

- **feed**
  - `feed-api`=8890
  - `feed-service`=10002
  - `feed-service-health`=20002
  - `feed-api-health`=30002
  - `depend`=[feedback,user]

- **publish**
  - `publish-api`=8891
  - `publish-service`=10003
  - `publish-service-health`=20003
  - `publish-api-health`=30003
  - `depend`=[user,feedback,]

- **favorite**
  - `favorite-api`=8892
  - `favorite-service`=10004
  - `favorite-service-health`=20004
  - `favorite-api-health`=30004
  - `depend`=[user,auths,feed,feedback]

- **comment**
  - `comment-api`=8893
  - `comment-service`=10005
  - `comment-service-health`=20005
  - `comment-api-health`=30005
  - `depend`=[auths,feedback]

- **relation**
  - `relation-api`=8894
  - `relation-service`=10006
  - `relation-service-health`=20006
  - `relation-api-health`=30006
  - `depend`=[user,auths]

- **feedback**
  - `feedback-api`=8895
  - `feedback-service`=10007
  - `feedback-service-health`=20007
  - `feedback-api-health`=30007

- **message**
  - `message-api`=8896
  - `message-service`=10008
  - `message-service-health`=20008
  - `message-api-health`=30008
  - `depend`=[relation,auths]