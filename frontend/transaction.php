<!DOCTYPE html>
<html lang="en">


<?php include "func/function.php"; ?>

<?php include "func/head.php"; ?>


<body class="d-flex flex-column min-vh-100 text-light" style="background-color: #111111;">

  <?php include "func/navbar.php"; ?>

  <br>
  <div class="container justify-content-space-between">
    <h4>Transactions</h4> <br>

    <div class="row" style="width: 100%;">

      <div class="col">
        <h5>Make a transaction</h5>
        <br>
        <p>Make a transaction as an admin. Choose processes and the amount of NBC to be transferred:</p>
        <br>
        <div>
          <form action="<?php echo htmlspecialchars($_SERVER["PHP_SELF"]); ?>" method="POST">
            <div class="mb-3" style="width: 55%;">

              <label for="sender" class="form-label">Sender</label>
              <input class="form-control" id="sender">
            </div>
            <div class="mb-3" style="width: 55%;">
              <label for="receiver" class="form-label">Receiver</label>
              <input class="form-control" id="receiver">
            </div>
            <div class="mb-3" style="width: 55%;">
              <label for="amount" class="form-label">Amount</label>
              <input class="form-control" id="amount">
            </div>
            <br>
            <div class="container d-flex justify-content-center">
              <button class="btn btn-light" type="submit">Send</button>
            </div>
          </form>
        </div>

      </div>
      <div class="col card d-flex " style="background-color: #111111;">
        <div class="card-header d-flex justify-content-left">
          <h5>Wallet balance</h5>
        </div>
        <div class="card-body d-flex justify-content-start text-white">
          <p><b style="font-size: 10vw;">55</b>
            <b style="font-size: 5vw;">NBC</b>
            <br> History
          </p>
        </div>
        <div class="overflow-auto" style="max-width: 100%; max-height: 200px; background-color: #222222">
          <table class="table  text-light" style="background-color: #222222">
            <thead>
              <tr>
                <th scope="col">#</th>
                <th scope="col">First</th>
                <th scope="col">Last</th>
                <th scope="col">Handle</th>
              </tr>
            </thead>
            <tbody>
              <tr>
                <th scope="row">1</th>
                <td>Mark</td>
                <td>Otto</td>
                <td>@mdo</td>
              </tr>
              <tr>
                <th scope="row">2</th>
                <td>Jacob</td>
                <td>Thornton</td>
                <td>@fat</td>
              </tr>
              <tr>
                <th scope="row">3</th>
                <td>Larry</td>
                <td>the Bird</td>
                <td>@twitter</td>
              </tr>
              <tr>
                <th scope="row">1</th>
                <td>Mark</td>
                <td>Otto</td>
                <td>@mdo</td>
              </tr>
              <tr>
                <th scope="row">2</th>
                <td>Jacob</td>
                <td>Thornton</td>
                <td>@fat</td>
              </tr>
              <tr>
                <th scope="row">3</th>
                <td>Larry</td>
                <td>the Bird</td>
                <td>@twitter</td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>



</body>


<?php include "func/footer.php"; ?>

</html>