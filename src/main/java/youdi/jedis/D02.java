package youdi.jedis;

import redis.clients.jedis.Jedis;
import redis.clients.jedis.JedisPool;
import redis.clients.jedis.Transaction;

import java.util.List;

public class D02 {
    public static void main(String[] args) {
        JedisPool instance = JedisPoolUtils.getJedisPoolInstance();


        Jedis jedis = instance.getResource();

        // 乐观锁
        jedis.watch("a");

        Transaction multi = jedis.multi();

        // 组队

        multi.decr("a");
        multi.sadd("b", "v");

        List<Object> exec = multi.exec();


        if (exec == null || exec.size() == 0) {
            System.out.println("秒杀失败");
            jedis.close();
        }

    }
}
