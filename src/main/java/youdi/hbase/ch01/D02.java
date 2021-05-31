package youdi.hbase.ch01;


import org.apache.hadoop.conf.Configuration;
import org.apache.hadoop.hbase.HBaseConfiguration;
import org.apache.hadoop.hbase.TableName;
import org.apache.hadoop.hbase.client.Admin;
import org.apache.hadoop.hbase.client.Connection;
import org.apache.hadoop.hbase.client.ConnectionFactory;
import org.apache.hadoop.hbase.client.Table;

import java.io.IOException;

public class D02 {
    public static void main(String[] args) throws IOException {
        Configuration conf = HBaseConfiguration.create();
        Connection conn = ConnectionFactory.createConnection(conf);


        TableName tableName = TableName.valueOf("youdi:student");

        Admin admin = conn.getAdmin();

        if(admin.tableExists(tableName)) {
//            Table table = conn.getTable(tableName);
            // 禁用表
//            admin.disableTable(tableName);

            // 删除表
//            admin.deleteTable(tableName);





        }


    }
}
