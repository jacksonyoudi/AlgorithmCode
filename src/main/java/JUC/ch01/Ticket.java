package JUC.ch01;

import java.util.concurrent.locks.ReentrantLock;

public class Ticket {

    private int number = 30;
    private final ReentrantLock lock = new ReentrantLock();

    public void sale() {
        this.lock.lock();
        try {
            if (number > 0) {
                number -= 1;
            }
        } finally {
            this.lock.unlock();
        }
    }

}
