package stream.ch01;

import java.util.stream.Stream;

/**
 * @program: AlgorithmCode
 * @description:
 * @author: changyouliang
 * @date: 2021/11/12
 **/
public class D01 {
    public static void main(String[] args) {
        Stream.of("apple", "banana", "orange", "waltermaleon", "grape")
                .map(String::length)
                .forEach(e ->System.out.println(e));


        Stream.of("apple", "banana", "orange", "waltermaleon", "grape")
                .mapToInt(String::length)
                .forEach(e ->System.out.println(e));
    }
}
