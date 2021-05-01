package youdi.hbase.util;

import org.apache.hadoop.conf.Configuration;
import org.apache.hadoop.hbase.HBaseConfiguration;
import org.apache.hadoop.hbase.TableName;
import org.apache.hadoop.hbase.client.Connection;
import org.apache.hadoop.hbase.client.ConnectionFactory;
import org.apache.hadoop.hbase.client.Put;
import org.apache.hadoop.hbase.client.Table;
import org.apache.hadoop.hbase.util.Bytes;


public class HbaseUtil {
    private static ThreadLocal<Connection> connHolder = new ThreadLocal<Connection>();


    public HbaseUtil() {
    }

    public static void makeHbaseConnection() throws Exception {
        Connection conn = connHolder.get();

        if (null == conn) {
            Configuration conf = HBaseConfiguration.create();
            conn = ConnectionFactory.createConnection(conf);
            connHolder.set(conn);
        }

    }


    public static void insertData(String tableName, String rowkey, String family, String colname, String value) throws Exception {
        Connection conn = connHolder.get();
        Table table = conn.getTable(TableName.valueOf(tableName));

        Put put = new Put(Bytes.toBytes(rowkey));
        put.addColumn(Bytes.toBytes(family),Bytes.toBytes(colname),Bytes.toBytes(value));

        table.put(put);
        table.close();

    }



    public static void close() throws Exception {
        Connection conn = connHolder.get();
        if (null != conn) {
            conn.close();
            connHolder.remove();
        }

    }


}
