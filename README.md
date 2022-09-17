# protobjc

跨平台序列化框架protobj的编译器，使用protobuf序列化格式，扩展了一些语法。


## example
### message
```protobuf
package example;
//导入一个消息
import other.MessageDef;

//message name[index]（index：拥有同一个父类的message所形成的一个组内id）
message ExampleMessage[1]{
  //field
  //modifier type name = number[option=value];
        string s = 1;
        i32   i = 2;
    lst string lsts = 3;
    set string sets = 4;
    ext OtherMessage other = 5;
    lst OtherMessage others = 6[polymorphic = true];
}
message OtherMessage[0]{
    ExampleEnum e = 1;
}
enum ExampleEnum{
  //name = number 
    e1 = 0;
    e2 = 1;
}
```

modifier：修饰符

* dft(default):默认
* arr：数组
* ext：扩展,需要与index配合使用：针对不同语言可以有不同的实现，java中用继承表示，没有继承的语言可以用当作一个字段
* lst：列表
* set：不可重复列表

type
* bool：true or false
* i8
* u8
* i16
* u16
* i32
* u32
* s32
* f32
* sf32
* i64
* u64
* s64
* f64
* sf64
* string
* double
* float
* map
* enum(custom)
* message(custom)
