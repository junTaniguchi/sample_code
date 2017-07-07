package com.xcat.aplus.api;

import java.io.File;
import java.io.IOException;
import java.util.ArrayList;

public class SearchFile {

	private String rootPath;

	public SearchFile(){
	}

	public SearchFile(String rootPath){
		this.rootPath = rootPath;
	}

	// ファイルを再帰処理で検索
	public ArrayList<File> recursionSearchFileName(ArrayList<File> hitFileList, String targetFileName) throws IOException{
		System.out.println("ファイル名検索");
		File searchDir = new File(this.rootPath);

		File objlist[] = searchDir.listFiles();

		for (File obj : objlist){
			if (obj.isDirectory()){
				//取得したオブジェクトがディレクトリの場合、再帰処理を行う
				this.rootPath = obj.getCanonicalPath();
				hitFileList = recursionSearchFileName(hitFileList, targetFileName);
			}else if (obj.isFile()){
				System.out.println(targetFileName);
				System.out.println(obj.getName());
				//取得したオブジェクトがファイルの場合、ファイル名が一致するかを確認する。
				if (targetFileName.equals(obj.getName())){
					hitFileList.add(obj);
				}
			}
		}
		return hitFileList;
	}
	// ファイルを再帰処理で検索
	public ArrayList<File> recursionSearchTargetDate(ArrayList<File> hitFileList, String targetDate) throws IOException{
		System.out.println("更新日付検索");

		File searchDir = new File(this.rootPath);

		File objlist[] = searchDir.listFiles();

		for (File obj : objlist){
			if (obj.isDirectory()){
				//取得したオブジェクトがディレクトリの場合、再帰処理を行う
				if (obj.getCanonicalPath().indexOf(targetDate) > -1){
					this.rootPath = obj.getCanonicalPath();
					hitFileList = recursionSearchTargetDate(hitFileList, targetDate);
				}
			}else if (obj.isFile()){
				//取得したオブジェクトがファイルの場合、ファイル名が一致するかを確認する。
				if (obj.getCanonicalPath().indexOf(targetDate) > -1){
					hitFileList.add(obj);
				}
			}
		}
		return hitFileList;
	}
	// ファイル名と更新日時を再帰処理で検索
	public ArrayList<File> recursionSearchFile(ArrayList<File> hitFileList, String targetFileName, String targetDate) throws IOException{
		System.out.println("2種類検索");

		File searchDir = new File(this.rootPath);

		File objlist[] = searchDir.listFiles();

		for (File obj : objlist){
			if (obj.isDirectory()){
				//取得したオブジェクトがディレクトリの場合、再帰処理を行う
				this.rootPath = obj.getCanonicalPath();
				hitFileList = recursionSearchFile(hitFileList, targetFileName, targetDate);
			}else if (obj.isFile()){
				System.out.println("File : " + obj.getName());
				System.out.println("targetFileName : " + targetFileName);
				//取得したオブジェクトがファイルの場合、ファイル名が一致するかを確認する。
				if (targetFileName.equals(obj.getName())){
					System.out.println("True");
					if (obj.getCanonicalPath().indexOf(targetDate) > -1){
						hitFileList.add(obj);
					}
				}else {
					System.out.println("False");
				}
			}
		}
		return hitFileList;
	}
}

