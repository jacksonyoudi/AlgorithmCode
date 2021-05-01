package youdi.hbase.ch01;

import org.apache.hadoop.conf.Configuration;
import org.apache.hadoop.hbase.*;
import org.apache.hadoop.hbase.client.*;
import org.apache.hadoop.hbase.util.Bytes;


import java.io.IOException;
import java.nio.charset.StandardCharsets;

public class D01 {
    public static void main(String[] args) throws IOException {


        Configuration conf = HBaseConfiguration.create();

        // classpath: hbase-default.xml  hbase-site.xml


        Connection conn = ConnectionFactory.createConnection(conf);


        // 操作对象 admin
        Admin admin = conn.getAdmin();

        try {
            admin.getNamespaceDescriptor("youdi");
        } catch (NamespaceNotFoundException e) {
            // 创建表空间
            NamespaceDescriptor youdi = NamespaceDescriptor.create("youdi").build();
            admin.createNamespace(youdi);
        }

        // 判断是否存在某种表
        TableName student = TableName.valueOf("youdi:student");
        boolean b = admin.tableExists(student);

        if (!b) {
            ColumnFamilyDescriptor info = ColumnFamilyDescriptorBuilder.of("info");

            TableDescriptor descriptor = TableDescriptorBuilder.newBuilder(student).setColumnFamily(info).build();

            admin.createTable(descriptor);

        } else {
            // 查询数据 ddl dml

            // 获取表对象
            Table table = conn.getTable(student);

            String rowkey = "1001";

//            byte[] bytes = rowkey.getBytes(StandardCharsets.UTF_8);
            Get get = new Get(Bytes.toBytes(rowkey));

            Result result = table.get(get);
            boolean empty = result.isEmpty();


            System.out.println(!empty);

            if (empty) {
                Put put = new Put(Bytes.toBytes(rowkey));
                String cf = "info";
                put.addColumn(Bytes.toBytes(cf), Bytes.toBytes("name"), Bytes.toBytes("jackson"));


                table.put(put);
                System.out.println("success");

            } else {
                for (Cell cell : result.rawCells()) {
                    System.out.println(Bytes.toString(CellUtil.cloneFamily(cell)));
                }
            }

        }


    }
}
