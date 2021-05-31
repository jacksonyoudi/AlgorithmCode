package youdi.hbase.mr;


import org.apache.hadoop.util.ToolRunner;
import youdi.hbase.tool.HbaseMapperReduceTool;

public class DTable {
    public static void main(String[] args) throws Exception {

        ToolRunner.run(new HbaseMapperReduceTool(), args);

    }
}
