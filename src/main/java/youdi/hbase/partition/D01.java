package youdi.hbase.partition;

import org.apache.hadoop.conf.Configuration;
import org.apache.hadoop.hbase.HBaseConfiguration;
import org.apache.hadoop.hbase.TableName;
import org.apache.hadoop.hbase.client.*;
import org.apache.hadoop.hbase.filter.BinaryComparator;
import org.apache.hadoop.hbase.filter.CompareFilter;
import org.apache.hadoop.hbase.filter.RowFilter;
import org.apache.hadoop.hbase.util.Bytes;

import java.io.IOException;

public class D01 {
    public static void main(String[] args) throws IOException {
        Configuration conf = HBaseConfiguration.create();

        // classpath: hbase-default.xml  hbase-site.xml


        Connection conn = ConnectionFactory.createConnection(conf);

        // 操作对象 admin
        Admin admin = conn.getAdmin();


        // 操作对象 admin
//        Admin admin = conn.getAdmin();

        TableName student = TableName.valueOf("student");
        ColumnFamilyDescriptor info = ColumnFamilyDescriptorBuilder.of("info");

        TableDescriptor descriptor = TableDescriptorBuilder.newBuilder(student).setColumnFamily(info).build();



        admin.createTable(descriptor);

    }
}
