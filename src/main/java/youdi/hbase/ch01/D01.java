package youdi.hbase.ch01;

import org.apache.hadoop.conf.Configuration;
import org.apache.hadoop.hbase.NamespaceDescriptor;
import org.apache.hadoop.hbase.NamespaceNotFoundException;
import org.apache.hadoop.hbase.client.Admin;
import org.apache.hadoop.hbase.client.Connection;
import org.apache.hadoop.hbase.client.ConnectionFactory;
import org.apache.hadoop.hbase.HBaseConfiguration;
import org.apache.hadoop.hbase.TableName;


import java.io.IOException;

public class D01 {
    public static void main(String[] args) throws IOException {


        Configuration conf = HBaseConfiguration.create();

        // classpath: hbase-default.xml  hbase-site.xml

        Connection conn = ConnectionFactory.createConnection(conf);


        // 操作对象 admin
        Admin admin = conn.getAdmin();

        // 判断是否存在某种表
        TableName student = TableName.valueOf("student");
        boolean b = admin.tableExists(student);
        System.out.println(b);

        try {
            admin.getNamespaceDescriptor("hbase");
        } catch (NamespaceNotFoundException e) {
            // 创建表空间


            admin.createNamespace();
        }


    }
}
