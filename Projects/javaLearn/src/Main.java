import java.sql.*;

class A {
  public void fpp() {
    System.out.println("foo method called");
  }

  static class B {
    public void bar() {
      System.out.println("Bar method called");
    }

  }
}

public class Main {

  public static void main(String[] args) {
    String url = "jdbc:mysql://localhost:3306/snippetbox";
    String user = "anurag";
    String password = "anurag616";

      try(Connection conn = DriverManager.getConnection(url,user,password)) {
        Snippets s = new Snippets(conn);
        s.getSnippets();
      } catch (SQLException e) {
          throw new RuntimeException(e);
      }

  }

}

class Snippets {
  Connection con;
  public Snippets(Connection con) {
    this.con = con;
  }

  public void getSnippets() throws SQLException{
    Statement st = con.createStatement();
    String qry = "SELECT * FROM snippets";
    ResultSet rs = st.executeQuery(qry);

    System.out.println("ID" + " " + "TITLE");
    while (rs.next()) {
     String title = rs.getString("title");
     int id = rs.getInt("id");
     System.out.println(id + "  " + title);
    }

  }
}
