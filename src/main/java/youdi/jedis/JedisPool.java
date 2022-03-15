package youdi.jedis;

import redis.clients.jedis.Jedis;
import redis.clients.jedis.JedisPool;
import redis.clients.jedis.JedisPoolConfig;

class JedisPoolUtils {
    private static volatile JedisPool jedisPool = null;

    public JedisPoolUtils() {
    }

    // 单例模式
    public static JedisPool getJedisPoolInstance() {
        if (null == jedisPool) {
            synchronized (JedisPoolUtils.class) {
                JedisPoolConfig config = new JedisPoolConfig();
                config.setMaxIdle(32);
                config.setMaxTotal(200);
                config.setMaxWaitMillis(100 * 1000);
                config.setBlockWhenExhausted(true);
                config.setTestOnBorrow(true);

                jedisPool = new JedisPool(config, "111", 6379, 5000);
            }
        }
        return jedisPool;
    }

    public static void release(JedisPool jedisPool, Jedis jedis) {
        if (null != jedis) {
//            jedisPool.getResource().r
        }
    }

}
