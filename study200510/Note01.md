# 04 패키지 (Packages)

Go는 디렉토리 경로 마지막 이름이 패키지 이름이다.

import를 여러 개 사용할 때 두가지 방식이 있다

# 05 임포트 (Imports)
```go
import (
  "fmt"
  "math"
)
```
or
```go
import "fmt"
import "math"
```
----------------------------------------------------------------------
**Nextafter(x double, y double) 메소드**

math 패키지 안의 export된 메소드로

x>y 인 경우 표현 가능한 x보다 작은 실수 중 제일 큰 실수를 return

x<y 인 경우 표현 가능한 x보다 큰 실수 중 제일 작은 실수를 return


Nextafterf(), Nextafterl()을 이용해 float와 long형 실수도 가능.

----------------------------------------------------------------------

# 06 익스포트 (Exported names)

패키지 import 시 패키지에서 외부로 export한 것들(메소드, 변수, 상수 등)에 접근가능

대문자로 시작하는 것들이 export 된 것들이다.


