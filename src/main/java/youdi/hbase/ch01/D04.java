package youdi.hbase.ch01;

import org.apache.hadoop.conf.Configuration;
import org.apache.hadoop.hbase.HBaseConfiguration;
import org.apache.hadoop.hbase.TableName;
import org.apache.hadoop.hbase.client.*;
import org.apache.hadoop.hbase.filter.BinaryComparator;
import org.apache.hadoop.hbase.filter.CompareFilter;
import org.apache.hadoop.hbase.util.Bytes;
import youdi.hbase.util.HbaseUtil;
import org.apache.hadoop.hbase.filter.RowFilter;


public class D04 {
    public static void main(String[] args) throws Exception {
        Configuration conf = HBaseConfiguration.create();

        // classpath: hbase-default.xml  hbase-site.xml


        Connection conn = ConnectionFactory.createConnection(conf);


        // 操作对象 admin
//        Admin admin = conn.getAdmin();

        TableName student = TableName.valueOf("student");
        Table table = conn.getTable(student);

        Scan scan = new Scan();

        BinaryComparator comparator = new BinaryComparator(Bytes.toBytes("2001"));

        RowFilter rowFilter = new RowFilter(CompareFilter.CompareOp.EQUAL, comparator);
        scan.setFilter(rowFilter);
        ResultScanner results = table.getScanner(scan);


    }
}
