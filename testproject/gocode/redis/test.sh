set -ex

BRANCH="ml"
redisUrl="game.sanguo.bj"
declare -A REDIS_URL_MAP=(['ml']="${redisUrl}:9898" ['jp']="${redisUrl}:9900" ['ie']="${redisUrl}:9902" ['sk']="${redisUrl}:9901" ['tw']="${redisUrl}:9903")
mm=${REDIS_URL_MAP[${BRANCH}]}