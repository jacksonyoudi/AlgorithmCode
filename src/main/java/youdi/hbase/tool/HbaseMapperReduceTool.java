package youdi.hbase.tool;

import org.apache.hadoop.conf.Configuration;
import org.apache.hadoop.hbase.TableName;
import org.apache.hadoop.hbase.client.Put;
import org.apache.hadoop.hbase.client.Scan;
import org.apache.hadoop.hbase.io.ImmutableBytesWritable;
import org.apache.hadoop.hbase.mapreduce.TableMapReduceUtil;
import org.apache.hadoop.mapred.JobStatus;
import org.apache.hadoop.mapreduce.Job;
import org.apache.hadoop.util.Tool;


public class HbaseMapperReduceTool implements Tool {
    @Override
    public int run(String[] strings) throws Exception {
        // 作业
        Job job = Job.getInstance();
        job.setJarByClass(HbaseMapperReduceTool.class);

        TableName tableName = TableName.valueOf("youdi:");

        // mapper
        TableMapReduceUtil.initTableMapperJob(
                tableName,
                new Scan(),
                ScanDataMapper.class,
                ImmutableBytesWritable.class,
                Put.class,
                job
        );


        TableMapReduceUtil.initTableReducerJob("youdi:student", InsertDataReducer.class, job
        );


        boolean stat = job.waitForCompletion(true);
        return stat ? JobStatus.State.SUCCEEDED.getValue() : JobStatus.State.FAILED.getValue();

    }

    @Override
    public void setConf(Configuration configuration) {

    }

    @Override
    public Configuration getConf() {
        return null;
    }
}
