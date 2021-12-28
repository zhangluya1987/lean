# 第二次作业说明
[toc]

## 助教老师辛苦了 以下是作业目录结构说明:
```go
1.第一个作业题代码目录
    - homework/01.fatRate_v2.0

2.第二个作业题代码目录
    - homework/02.fatRate_UnitTesting_v1.0
```


## 第一个作业
```go
1.扩展功能依赖码路径:
    - https://github.com/zhangluya1987/learn/tree/master/staging/src/github.com/armstrongli/go-bmi
2.项目外部核心功能依赖:
    - https://github.com/armstrongli/go-bmi

3.项目执行方式
[jesse@infra 01.fatRate_v2.0 %]# go run ./main.go --Name Jesse --Sex "男" --Height 1.75 --Weight 70 --Age 27
    Name: Jesse
    Sex: 男
    Height: 1.75
    Weight: 70
    Age: 27
    bmi:22.857142857142858,bfr:0.17438571428571426

4.第一个作业代码直达链接:
    - https://github.com/zhangluya1987/learn/tree/master/lesson3/homework/01.fatRate_v2.0
```

## 第二个作业
### (一)内容:为体脂计算器编写单元测试并完善体脂计算器的验证逻辑。
```go
1.第一部分UnitTest:
    // 要求
    BMI 计算：
    录入正常身高、体重，确保计算结果符合预期；
    录入 0 或负数身高，返回错误；
    录入 0 或负数体重，返回错误。

    // 关联代码位置
    https://github.com/zhangluya1987/learn/tree/master/lesson3/homework/02.fatRate_UnitTesting_v1.0/bmi

    // 执行单测方式
    [jesse@infra bmi %]# go test -v .
    === RUN   TestBMICase1
    --- PASS: TestBMICase1 (0.00s)
    === RUN   TestBMICase2
        bmi_test.go:29: 身高数值输入异常: 身高不能为0。
        bmi_test.go:37: 身高数值输入异常: 身高不能为负数。
    --- FAIL: TestBMICase2 (0.00s)
    === RUN   TestBMICase3
        bmi_test.go:48: 体重数值输入异常: 体重不能为0。
        bmi_test.go:56: 体重数值输入异常: 体重不能为负数。
    --- FAIL: TestBMICase3 (0.00s)
    FAIL
    FAIL    lesson3/homework/02.fatRate_UnitTesting_v1.0/bmi        0.418s
    FAIL


2.第二部分UnitTest:
    // 要求
    体脂率计算：
    录入正常 BMI、年龄、性别，确保计算结果符合预期；
    录入非法 BMI、年龄、性别（0、负数、超过 150 的年龄、非男女的性别输入），返回错误；
    录入完整的性别、年龄、身高、体重，确保最终获得的健康建议符合预期。

    //关联代码位置
    https://github.com/zhangluya1987/learn/tree/master/lesson3/homework/02.fatRate_UnitTesting_v1.0/bfr

    // 引用代码部分
    https://github.com/zhangluya1987/learn/tree/master/lesson3/homework/02.fatRate_UnitTesting_v1.0/suggest

    // 执行单测方式
    [jesse@infra bfr %]# go test -v .
    === RUN   TestCalcFatRateCase1
    --- PASS: TestCalcFatRateCase1 (0.00s)
    === RUN   TestCalcFatRateCase2
        fatrate_test.go:39: 输入值bmi=0,但期望输入不能为0。
        fatrate_test.go:49: bmi=-1,但期望输入不能为负数。
        fatrate_test.go:59: 输入年龄为:160,但期望输入不能大于150岁。
        fatrate_test.go:69: 输入年龄为:-30,但期望输入不能小于18岁。
        fatrate_test.go:79: 输入的性别为:我就不输入女或者男,不合法，请输入男或者女
            
    --- FAIL: TestCalcFatRateCase2 (0.00s)
    === RUN   TestHealthSuggestCase3
    --- PASS: TestHealthSuggestCase3 (0.00s)
    FAIL
    FAIL    lesson3/homework/02.fatRate_UnitTesting_v1.0/bfr        0.401s
    FAIL

    Created by Jesse
        -- 20211228
```