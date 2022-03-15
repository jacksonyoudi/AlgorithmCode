package youdi.hbase.coprocesser;

import org.apache.hadoop.hbase.TableName;
import org.apache.hadoop.hbase.client.Put;
import org.apache.hadoop.hbase.client.Table;
import org.apache.hadoop.hbase.coprocessor.ObserverContext;
import org.apache.hadoop.hbase.coprocessor.RegionCoprocessorEnvironment;
import org.apache.hadoop.hbase.coprocessor.RegionObserver;
import org.apache.hadoop.hbase.wal.WALEdit;

import java.io.IOException;


/**
 *
 */
public class MyCoprocesser implements RegionObserver {
    @Override
    public void postPut(ObserverContext<RegionCoprocessorEnvironment> c, Put put, WALEdit edit) throws IOException {
        RegionCoprocessorEnvironment environment = c.getEnvironment();
        Table student = environment.getConnection().getTable(TableName.valueOf("student"));

        student.put(put);
        student.close();

    }
}
