package JUC.ch02.rw;


import java.util.HashMap;
import java.util.Map;
import java.util.concurrent.TimeUnit;
import java.util.concurrent.locks.ReadWriteLock;
import java.util.concurrent.locks.ReentrantReadWriteLock;

class MyCache {
    private ReadWriteLock lock = new ReentrantReadWriteLock();


    private volatile Map<String, Object> map = new HashMap<String, Object>();


    public void put(String key, Object value) {

        lock.writeLock().lock();

        try {
            System.out.println(Thread.currentThread().getName() + " 正在写操作" + key);


            try {
                TimeUnit.MICROSECONDS.sleep(300);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }

            map.put(key, value);
            System.out.println(Thread.currentThread().getName() + " 写完了");
        } finally {
            lock.writeLock().unlock();
        }
    }


    public Object get(String key) {
        lock.readLock().lock();

        Object result = null;
        try {

            System.out.println(Thread.currentThread().getName() + " 正在读取" + key);

            TimeUnit.MICROSECONDS.sleep(300);

             result = map.get(key);
            System.out.println(Thread.currentThread().getName() + " 读完了");

            return result;
        } catch (Exception e) {
            e.printStackTrace();

        } finally {
            lock.readLock().unlock();
        }
        return result;
    }
}


public class D01 {
    public static void main(String[] args) {
        MyCache cache = new MyCache();

    }
}
