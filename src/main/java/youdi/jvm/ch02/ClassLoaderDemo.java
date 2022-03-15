package youdi.jvm.ch02;


/**
 * sun.misc.Launcher$AppClassLoader@18b4aac2
 * sun.misc.Launcher$ExtClassLoader@1cf4f579
 * null
 * sun.misc.Launcher$AppClassLoader@18b4aac
 */
public class ClassLoaderDemo {
    public static void main(String[] args) {
        ClassLoader loader = ClassLoader.getSystemClassLoader();
        System.out.println(loader);  // sun.misc.Launcher$AppClassLoader@18b4aac2

        // 获取上层 扩展类加载器
        ClassLoader parent = loader.getParent();
        System.out.println(parent);

        // 获取不到 引导类加载器的父类
        ClassLoader parent1 = parent.getParent();
        System.out.println(parent1); // null


        // String类使用引导类加载器进行加载的，---> Java的核心类库都是使用引导类加载器进行加载的。
        ClassLoader classLoader = ClassLoaderDemo.class.getClassLoader();
        System.out.println(classLoader); // sun.misc.Launcher$AppClassLoader@18b4aac2


    }
}
