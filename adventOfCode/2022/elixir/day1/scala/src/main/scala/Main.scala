import scala.io.Source
import scala.collection.mutable.ListBuffer
object DayOne {
  def main(args: Array[String]): Unit = {
    val filename: String = "input"

    val linesIterator = readFile(filename)

    var maxCalNumber: ListBuffer[Int] = ListBuffer()
    var currentElfCal = 0;

    for (line <- linesIterator) {
      if (line.length() == 0) {
        maxCalNumber += currentElfCal;
        currentElfCal = 0;
      } else {
        currentElfCal += Integer.parseInt(line);
      }
    }
    println(maxCalNumber.toList.sortWith(_ > _).take(3).sum)
  }

  private def readFile(filename: String): Iterator[String] = {
    Source.fromFile(filename).getLines()
  }
}
