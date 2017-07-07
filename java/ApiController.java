package com.xcat.aplus.api;

import static spark.Spark.*;

import java.io.BufferedReader;
import java.io.File;
import java.io.FileInputStream;
import java.io.InputStreamReader;
import java.text.SimpleDateFormat;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.Map;

import com.fasterxml.jackson.databind.ObjectMapper;

public class ApiController {

	static final ObjectMapper OM = new ObjectMapper();

	public static void main(String[] args) {
		// 環境変数のクラスパスへ下記の設定を行う。
		// 「APLUS_API\」ディレクトリまでの絶対パス
		staticFileLocation("/html");

		get("/stop", (req, res) -> {
			stop();
			return null;
		});
		get("/filepath", (req, res) -> {
			System.out.println("filepath");

			String rootPath = "C:\\Users\\j13-taniguchi\\Desktop\\EP";
			String targetFileName = req.queryParams("fileName");
			String targetDate = req.queryParams("fileDate");

			System.out.println("rootPath : " + rootPath);
			System.out.println("targetFileName : " + targetFileName);
			System.out.println("targetDate : " + targetDate);


			//クエリパラメータを元にオブジェクトを生成
			SearchFile searchFile = new SearchFile(rootPath);

			//クエリパラメータのパターンにて検索パターンを変更
			ArrayList<File> hitFileList = new ArrayList<File>();
			if (targetFileName != null && targetDate != null) {
				hitFileList = searchFile.recursionSearchFile(hitFileList, targetFileName, targetDate);
			} else if (targetFileName == null){
				hitFileList = searchFile.recursionSearchTargetDate(hitFileList, targetDate);
			} else {
				hitFileList = searchFile.recursionSearchFileName(hitFileList, targetFileName);
			}

			ArrayList<Object> fileConfig = new ArrayList<Object>();

			for (File file : hitFileList) {
				Map<String, Object> map = new HashMap<>();

				//ファイルの名前
				String filename = file.getName();

				//ファイルの更新日時
				SimpleDateFormat sdf = new SimpleDateFormat("yyyy/MM/dd HH:mm:ss");
				Long lastModified = file.lastModified();
				String filedate = sdf.format(lastModified);

				//ファイルの格納されている絶対パスを取得
				String filePath;
				filePath = file.getAbsolutePath();

				//ファイルの内容を文字列化
				String filetext;
				BufferedReader br = null;
				try {
					br = new BufferedReader(new InputStreamReader(new FileInputStream(file)));
					StringBuffer sb = new StringBuffer();
					int c;
					while ((c = br.read()) != -1) {
						sb.append((char) c);
					}
					filetext = sb.toString();
				}finally {
					br.close();
				}

				map.put("filename", filename);
				map.put("filedate", filedate);
				map.put("filePath", filePath);
				map.put("filetext", filetext);

				fileConfig.add(map);

			}

			//Javaオブジェクト→JSON文字列
			res.type("application/json");
			String jsonString = OM.writeValueAsString(fileConfig);

			System.out.println(jsonString);

			return jsonString;

		});
	}

}
