# Consume-Repair

## 📘 About Consume-Repair

Consume-Repair คือ แอปพลิเคชั่นที่ใช้สำหรับดึงข้อมูลจาก kafka เพื่อรันซ้พของระบบ A-Reward และบันทึกข้อมูลลงฐานข้อมูล สำหรับ sattlement point

## 📦 Built With

- [x] GO 1.13
- [x] Viper 1.4.0

## 🏷 Version
version 0.1.0

- [X] ดึงข้อมูลจาก kafka โดยระบุวันที่เริ่มต้นและสิ้นของ ของ transection ที่ต้องการ

## ⚙ How to use
1. Clone project

2. Set up environment

    ``` bash
    export ENV=dev
    ```

3. Run project by command

    ``` bash
    go run main.go
    ```