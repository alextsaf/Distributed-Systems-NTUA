<!DOCTYPE html>
<html lang="en" data-bs-theme="dark">

<?php include "func/function.php"; ?>


<?php include "func/head.php"; ?>


<body class="d-flex flex-column min-vh-100 text-light" style="background-color: #111111;">

  <?php include "func/navbar.php"; ?>



  <br><br>
  <div class="container-md">
    <h4>Admin Endpoints</h4>
    <br>
    <p>Here lies a way control everything:</p>
    <br>
    <form action="<?php echo htmlspecialchars($_SERVER["PHP_SELF"]); ?>" method="POST">
      <div class="row">
        <div class="col text-center">
          <button type="submit" class="btn btn-light " name="Endpoint" value="initialization">Initialize Nodes</button>
        </div>
        <div class="col text-center ">
          <button type="submit" class="btn btn-light" name="Endpoint" value="termination">Terminate all Nodes</button>
        </div>
        <div class="col text-center">
          <button type="submit" class="btn btn-light" name="Endpoint" value="run">Run Transaction Script</button>
        </div>
      </div>
    </form>
  </div>
  <?php
  $host = "TO DO";
  if ($_SERVER["REQUEST_METHOD"] == 'POST') {
    if ($_POST["Endpoint"] == "initialization") {
      $response = sendrequest('http://' . $host . ':endpoint gia init', 'POST');
    } else 
    if ($_POST["Endpoint"] == "termination") {
      $endpoint = $_POST["Endpoint"];
      $response = sendrequest('http://' . $host . ':endpoint gia termination' . $endpoint, 'POST');
    } else {
      $endpoint = $_POST["Endpoint"];
      $response = sendrequest('http://' . $host . ':endpoint gia termination' . $endpoint, 'POST');
    }
    if ($response["code"] != 200) {
      $json = json_decode($response["response"]);
  ?>
      <br><br>
      <div class="container  card d-flex justify-content-center" style="background-color: #222222;">
        <div class="card-header d-flex justify-content-center">
          Response
        </div>
        <div class="card-body d-flex justify-content-center bg-danger text-white">
          Failed!
        </div>
      </div>
    <?php
    } else {
    ?>
      <br><br>
      <div class="container-fluid card d-flex justify-content-center">
        <div class="card-header d-flex justify-content-center">
          Response
        </div>
        <div class="card-body d-flex justify-content-center bg-success text-white">
          Success!
        </div>
      </div>
  <?php
    }
  }
  ?>



</body>

<?php include "func/footer.php"; ?>

</html>