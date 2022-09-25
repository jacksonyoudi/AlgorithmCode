package youdi.boolmfilter;

import java.io.Serializable;
import java.nio.charset.Charset;
import java.nio.charset.StandardCharsets;
import java.security.MessageDigest;
import java.security.NoSuchAlgorithmException;
import java.util.BitSet;
import java.util.Collection;

public class BloomFilter<E> implements Serializable {
    private BitSet bitset;
    private int bitSetSize; // bitset 's size
    private double bitsPerElement;
    private int expectedNumberOfFilterElements;  // 期望保存多少
    private int numberOfAddedElements; // 保存了多少数据
    private int k;   // 一个val占多少个bit


    static final Charset charset = StandardCharsets.UTF_8;

    static final String hashName = "MD5"; // MD5 gives good enough accuracy in most circumstances. Change to SHA1 if it's needed
    static final MessageDigest digestFunction;

    static {
        MessageDigest tmp;
        try {
            tmp = java.security.MessageDigest.getInstance(hashName);
        } catch (NoSuchAlgorithmException e) {
            tmp = null;
        }


        digestFunction = tmp;
    }


    public BloomFilter(double c, int n, int k) {
        this.expectedNumberOfFilterElements = c;
        this.k = k;
        this.bitsPerElement = c;
        //  k = 0.618 * m / n
        this.bitSetSize = (int) Math.ceil(c * n);
        this.numberOfAddedElements = 0;
        this.bitset = new BitSet(bitSetSize);

    }

    public BloomFilter(int bitSetSize, int expectedNumberOfFilterElements) {
        this.bitSetSize = bitSetSize;
        this.expectedNumberOfFilterElements = expectedNumberOfFilterElements;
    }

    public BloomFilter(double falsePositiveProbability, int expectedNumberOfElements) {
        this(Math.ceil(-(Math.log(falsePositiveProbability) / Math.log(2))) / Math.log(2), // c = k / ln(2)
                expectedNumberOfElements,
                (int) Math.ceil(-(Math.log(falsePositiveProbability) / Math.log(2))));
    }


    // having bitset
    public BloomFilter(int bitSetSize, int expectedNumberOfFilterElements, int actualNumberOfFilterElements, BitSet filterData) {
        this(bitSetSize, expectedNumberOfFilterElements);
        this.bitset = filterData;
        this.numberOfAddedElements = actualNumberOfFilterElements;
    }


    public static int createHash(String val) {
        return createHash(val, charset);
    }

    public static int createHash(String val, Charset charset) {
        return createHash(val.getBytes(charset));
    }

    public static int createHash(byte[] data) {
        return createHashes(data, 1)[0];
    }

    public static int[] createHashes(byte[] data, int hashes) {
        int[] result = new int[hashes];

        int k = 0;
        byte salt = 0;
        while (k < hashes) {
            byte[] digest;
            synchronized (digestFunction) {
                digestFunction.update(salt);
                salt++;
                digest = digestFunction.digest(data);
            }

            for (int i = 0; i < digest.length / 4 && k < hashes; i++) {
                int h = 0;
                for (int j = (i * 4); j < (i * 4) + 4; j++) {
                    h <<= 8;
                    h |= ((int) digest[j]) & 0xFF;
                }
                result[k] = h;
                k++;
            }
        }
        return result;
    }


    @Override
    public boolean equals(Object obj) {
        if (obj == null) {
            return false;
        }
        if (getClass() != obj.getClass()) {
            return false;
        }
        final BloomFilter<E> other = (BloomFilter<E>) obj;
        if (this.expectedNumberOfFilterElements != other.expectedNumberOfFilterElements) {
            return false;
        }
        if (this.k != other.k) {
            return false;
        }
        if (this.bitSetSize != other.bitSetSize) {
            return false;
        }
        if (this.bitset != other.bitset && (this.bitset == null || !this.bitset.equals(other.bitset))) {
            return false;
        }
        return true;
    }

    @Override
    public int hashCode() {
        int hash = 7;
        hash = 61 * hash + (this.bitset != null ? this.bitset.hashCode() : 0);
        hash = 61 * hash + this.expectedNumberOfFilterElements;
        hash = 61 * hash + this.bitSetSize;
        hash = 61 * hash + this.k;
        return hash;
    }


    public double expectedFalsePositiveProbability() {
        return getFalsePositiveProbability(expectedNumberOfFilterElements);
    }

    public double getFalsePositiveProbability(double numberOfElements) {
        // (1 - e^(-k * n / m)) ^ k
        return Math.pow((1 - Math.exp(-k * (double) numberOfElements
                / (double) bitSetSize)), k);

    }

    public double getFalsePositiveProbability() {
        return getFalsePositiveProbability(numberOfAddedElements);
    }


    public int getK() {
        return k;
    }

    public void clear() {
        bitset.clear();
        numberOfAddedElements = 0;
    }

    public void add(E element) {
        add(element.toString().getBytes(charset));
    }


    public void add(byte[] bytes) {
        // 得到对应的bit
        int[] hashes = createHashes(bytes, k);

        // 对bitset设置1
        for (int hash : hashes) {
            bitset.set(Math.abs(hash % bitSetSize), true);
        }
        numberOfAddedElements++;
    }

    public void addAll(Collection<? extends E> c) {
        for (E element : c) {
            add(element);
        }
    }

    public boolean contains(E element) {
        return contains(element.toString().getBytes(charset));
    }


    // 判断是否在bitset中
    public boolean contains(byte[] bytes) {
        // 对应data得到的 bit数组
        int[] hashes = createHashes(bytes, k);
        for (int hash : hashes) {
            // 判断对应位置是否是 1
            if (!bitset.get(Math.abs(hash % bitSetSize))) {
                return false;
            }
        }
        return true;
    }


    public boolean containsAll(Collection<? extends E> c) {
        for (E element : c) {
            if (!contains(element)) {
                return false;
            }
        }
        return true;
    }


    // 对应位置的值
    public boolean getBit(int bit) {
        return bitset.get(bit);
    }

    public void setBit(int bit, boolean value) {
        bitset.set(bit, value);
    }

    public BitSet getBitSet() {
        return bitset;
    }

    public int size() {
        return this.bitSetSize;
    }

    public int count() {
        return this.numberOfAddedElements;
    }


    public int getExpectedNumberOfElements() {
        return expectedNumberOfFilterElements;
    }


    public double getExpectedBitsPerElement() {
        return this.bitsPerElement;
    }

    // 空间使用率
    public double getBitsPerElement() {
        return this.bitSetSize / (double) numberOfAddedElements;
    }
}

