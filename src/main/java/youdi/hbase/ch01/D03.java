package youdi.hbase.ch01;

import youdi.hbase.util.HbaseUtil;

public class D03 {
    public static void main(String[] args) throws Exception {
        HbaseUtil.makeHbaseConnection();

        HbaseUtil.insertData("youdi:student", "1002", "info", "name", "rose");

        HbaseUtil.close();


    }
}
