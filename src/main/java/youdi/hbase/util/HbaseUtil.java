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
        put.addColumn(Bytes.toBytes(family), Bytes.toBytes(colname), Bytes.toBytes(value));

        table.put(put);
        table.close();

    }

    /**
     * 生成分区键
     * @param regionNum
     * @return
     */
    public static byte[][] genRegionKeys(int regionNum) {
        byte[][] bs = new byte[regionNum - 1][];

        for (int i = 0; i < regionNum-1; i++) {
            bs[i] = Bytes.toBytes(i + "|");
        }

        return bs;

    }


    /**
     * 生成分区号
     *
     * @param rowkey
     * @return
     */
    public static String genRegionNum(String rowkey, int regionCount) {
        int regionNum;
        int hashCode = rowkey.hashCode();
        if (regionCount > 0 && (regionCount & (regionCount - 1)) == 0) {
            // 2  n
            regionNum = hashCode & (regionCount - 1);
        } else {
            regionNum = hashCode % (regionCount - 1);
        }
        return regionNum + "_" + rowkey;
    }


    public static void close() throws Exception {
        Connection conn = connHolder.get();
        if (null != conn) {
            conn.close();
            connHolder.remove();
        }

    }


}
