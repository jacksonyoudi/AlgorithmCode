package youdi.jvm.ch02;

import sun.misc.Launcher;

import java.net.URL;


/**
 * file:/Library/Java/JavaVirtualMachines/adoptopenjdk-8.jdk/Contents/Home/jre/lib/resources.jar
 * file:/Library/Java/JavaVirtualMachines/adoptopenjdk-8.jdk/Contents/Home/jre/lib/rt.jar
 * file:/Library/Java/JavaVirtualMachines/adoptopenjdk-8.jdk/Contents/Home/jre/lib/sunrsasign.jar
 * file:/Library/Java/JavaVirtualMachines/adoptopenjdk-8.jdk/Contents/Home/jre/lib/jsse.jar
 * file:/Library/Java/JavaVirtualMachines/adoptopenjdk-8.jdk/Contents/Home/jre/lib/jce.jar
 * file:/Library/Java/JavaVirtualMachines/adoptopenjdk-8.jdk/Contents/Home/jre/lib/charsets.jar
 * file:/Library/Java/JavaVirtualMachines/adoptopenjdk-8.jdk/Contents/Home/jre/lib/jfr.jar
 * file:/Library/Java/JavaVirtualMachines/adoptopenjdk-8.jdk/Contents/Home/jre/classes
 */
public class ClassLoaderDemo2 {
    public static void main(String[] args) {
        System.out.println("**************启动类加载器*******************");

        URL[] urLs = Launcher.getBootstrapClassPath().getURLs();

        for (URL element : urLs) {
            System.out.println(element.toExternalForm());
        }

        System.out.println("****扩展类加载器**********");
        String s = System.getProperty("java.ext.dirs");

        for (String element : s.split(":")) {
            System.out.println(element);
        }

        // /Users/youdi/Library/Java/Extensions
        ///Library/Java/JavaVirtualMachines/adoptopenjdk-8.jdk/Contents/Home/jre/lib/ext
        ///Library/Java/Extensions
        ///Network/Library/Java/Extensions
        ///System/Library/Java/Extensions
        ///usr/lib/java
    }
}
