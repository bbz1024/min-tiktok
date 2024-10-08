package keys

const (
	ActorID = "actor_id"

	UserTokenKey = "user_token:%s" // set

	UserInfoKey   = "user_info:%d"   // hash
	FollowCount   = "follow_count"   // 关注数量
	FollowerCount = "follower_count" // 粉丝数量
	TotalFavorite = "total_favorite" // 获赞数量
	WorkCount     = "work_count"     // 作品数量
	FavoriteCount = "favorite_count" // 点赞数量

	UserWorkKey = "user_work:%d" // set

	UserFavoriteKey = "user_favorite:%d" // set

	UserFollow    = "user_follow:%d"
	UserFollower  = "user_follower:%d"
	UserFriendKey = "user_friend:%d"
	UserDeedKey   = "user_deed:%s" // zset

	// -------------------- local cache --------------------
	UserLocalCacheKey = "user_local_cache:%d"
)
