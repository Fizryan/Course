package com.method.student;

/**
 *
 * @author Hafizryandin HM
 */
public class Student {
    String name;
    String studentId;
    double bahasaNilai;
    double englishNilai;
    double mathNilai;
    
    public Student(String name, String studentId, double bahasaNilai, double englishNilai, double mathNilai) {
        this.name = name;
        this.studentId = studentId;
        this.bahasaNilai = bahasaNilai;
        this.englishNilai = englishNilai;
        this.mathNilai = mathNilai;
    }
    
    public double calculateAverage(){
        return ((bahasaNilai + englishNilai + mathNilai) /3);
    }
    
    public void printStudentInfo() {
        System.out.println("Name: " + name);
        System.out.println("Student Id: " + studentId);
        System.out.println("Nilai Bahasa: " + bahasaNilai);
        System.out.println("Nilai English: " + englishNilai);
        System.out.println("Nilai Matematika: " + mathNilai);
        System.out.println("Rata-Rata Nilai: " + calculateAverage());
    }
    
    public void changeStudentId(String newId) {
        this.studentId = newId;
        System.out.println("ID Telah Di Ubah Menjadi: " + studentId);
    }
    
    public static void main(String[] args) {
        Student student = new Student("Hafizryandin Haykal Matondang", "000000000000", 87.40, 97.21, 76.21);
        student.printStudentInfo();
        student.changeStudentId("103022300158");
        student.printStudentInfo();
    }
}
