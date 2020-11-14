package com.wechat;

import java.security.MessageDigest;


public class WebUin {


    /**
     * 根据imei和uin生成的md5码，获取数据库的密码（去前七位的小写字母）
     *
     * @param imei
     * @param uin
     * @return
     */
    public void initDbPassword(String imei, String uin) {
        String md5 = this.md5(imei + uin);
        String password = md5.substring(0, 7).toLowerCase();
        System.out.println("key:" + password);
    }


    /**
     * md5加密
     *
     * @param content
     * @return
     */
    private String md5(String content) {
        MessageDigest md5 = null;
        try {
            md5 = MessageDigest.getInstance("MD5");
            md5.update(content.getBytes("UTF-8"));
            byte[] encryption = md5.digest();//加密
            StringBuffer sb = new StringBuffer();
            for (int i = 0; i < encryption.length; i++) {
                if (Integer.toHexString(0xff & encryption[i]).length() == 1) {
                    sb.append("0").append(Integer.toHexString(0xff & encryption[i]));
                } else {
                    sb.append(Integer.toHexString(0xff & encryption[i]));
                }
            }
            return sb.toString();
        } catch (Exception e) {
            e.printStackTrace();
            return null;
        }
    }


    public static void main(String[] args) {
        String imei = "861446041708192"; // d23fee7 1167cd0
//        String imei = "863741047691964"; //  ca097db, a9969ac
//        String imei = "861446045105114";  //  b253666  df7fe75
        String uin = "1597598733";
        WebUin w = new WebUin();
        w.initDbPassword(imei, uin);
    }

}
