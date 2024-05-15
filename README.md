
Hexagonal Architecture คืออะไร ?
Hexagonal architecture ถูกคิดค้นโดย Alistair Cockburn เมื่อปี 2005 โดยเป็น 1 ใน Code architecture pattern ที่เป็นที่นิยมของ Go ที่สร้างอยู่บน 2 pattern ใหญ่ๆคือ Adaptor Pattern และ Dependency Injection โดย

Hexagonal จะนำเสนอส่วนที่เป็น "business logic" ขึ้นมาโดยไม่สนใจว่า ส่วนที่มาเชื่อมต่อจะเป็นส่วนไหน ขอแค่ทำตาม spec ของ business logic = สามารถใช้งานได้แล้ว
เป็นการแยกส่วนกันระหว่าง inside part (business logic, domain layer) และ outside part (ส่วนเชื่อมต่อกับภายนอกเช่น UI, database, external service - interface layer)
hexagonal-architecture

![alt text](https://miro.medium.com/v2/resize:fit:609/1*9j6W2a9wxj887QjDT8VyBQ.png)

Ref: https://www.tomkdickinson.co.uk/hexagonal-architecture-with-go-and-google-wire-e4344dd24b94

โดยองค์ประกอบใหญ่ๆของ Hexagonal Architecture จะประกอบด้วยของ 3 อย่างใหญ่ๆคือ

Ports ส่วน interfaces ที่บอกว่า application สามารถเชื่อมต่อมายัง business logic ได้ยังไง หรือ สามารถเข้าถึง resource ภายนอกได้ยังไง
Adapters ส่วนที่จะมาเชื่อมต่อกับ Port เป็นเหมือน "สะพาน" ระหว่าง Resource จริง กับ business logic (เช่น เป็นคนคุยกับ database, web service เป็นต้น)
Domain-centric ส่วนตรงกลางของ Business logic เป็นเหมือนศูนย์กลางของการคำนวน
hexagonal-port

![alt text](https://miro.medium.com/v2/resize:fit:1400/format:webp/1*ESpow8v_KPsCm2gdMUEsvw.png)

Ref: https://medium.com/@ebubekiryigit/hexagonal-architecture-a-golang-perspective-7eb3cb6117e7