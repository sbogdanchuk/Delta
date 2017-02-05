def delta(t1:String, t2:String):String={
	val tupA=parse(t1)
	val tupB=parse(t2)
	val res=tupA._1*60+tupA._2-tupB._1*60-tupB._2
	val resH=res/60
	return resH.toString+":"+("%02d".format(res-resH*60)).toString
}

def parse(s:String):Tuple2[Int,Int]=
{
	val arr=if(s.contains(":"))s.split(":") else s.split("-")
	return (if(arr(0).isEmpty)0 else arr(0).toInt, if(arr.length==1)0 else arr(1).toInt)
}

if(args.length >=2)
printf("Total=%s%n",args.reduceLeft((a,b)=>delta(a,b)))
else if (args.length==1){printf("Total=%s%n", args(0))
}

