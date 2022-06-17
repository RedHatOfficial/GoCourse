Compiled from "Test.java"
class Test {
  Test();
    Code:
       0: aload_0
       1: invokespecial #1          // Method java/lang/Object."<init>":()V
       4: return

  public static void main(java.lang.String[]);
    Code:
       0: new           #2          // class java/util/ArrayList
       3: dup
       4: invokespecial #3          // Method java/util/ArrayList."<init>":()V
       7: astore_1
       8: new           #2          // class java/util/ArrayList
      11: dup
      12: invokespecial #3          // Method java/util/ArrayList."<init>":()V
      15: astore_2
      16: aload_1
      17: bipush        42
      19: invokestatic  #4          // Method java/lang/Integer.valueOf:(I)Ljava/lang/Integer;
      22: invokeinterface #5,  2    // InterfaceMethod java/util/List.add:(Ljava/lang/Object;)Z
      27: pop
      28: aload_2
      29: ldc           #6          // String Hello
      31: invokeinterface #5,  2    // InterfaceMethod java/util/List.add:(Ljava/lang/Object;)Z
      36: pop
      37: return
}
