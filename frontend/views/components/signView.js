export class SignView {
  constructor() {
    this.sign = sign;
  }

  SignUpComponent() {
    const signUpCard = document.createElement("section");
    signUpCard.classList.add("sign-up-card");
    const signUpTemplate = `
<div class="sign-up-card">
<div class="sign-up-container">
<div class="sign-up-header">
<h2>Sign Up</h2>
</div>
<div class="sign-up-content">
<form method="post" class="sign-up-form" id="sign-up-form">
<div class="sign-up-form-group">
<label for="username">Username</label>
<input type="text" id="username" name="username" placeholder="Enter Your Username" required>
</div>
<div class="sign-up-form-group">
<label for="email">Email</label>
<input type="email" id="email" name="email" placeholder="Enter Your Email" required>
</div>
<div class="sign-up-form-group">
<label for="password">Password</label>
<input type="password" id="password" name="password" placeholder="Enter Your Password" required>
<span class="material-symbols-outlined visibility">visibility_off</span>
</div>
<div class="sign-up-form-group">
<label for="confirm-password">Confirm Password</label>
<input type="password" id="confirm-password" name="confirm-password" placeholder="Enter Your Password Again" required>
<span class="material-symbols-outlined visibility">visibility_off</span>
</div>
<div class="sign-up-form-group">
<label for="bio">Bio</label>
<input type="text" id="bio" name="bio" required>
</div>
<div class="sign-up-form-group">
<button type="submit">Sign Up</button>
</div>
</form>
</div>
</div>
</div>
`;
    signUpCard.innerHTML = signUpTemplate;
    return signUpCard;
  }

  SignInComponent() {
    const signInCard = document.createElement("section");
    signInCard.classList.add("sign-in-card");
    const signInTemplate = `
<div class="sign-in-card">
<div class="sign-in-container">
<div class="sign-in-header">
<h2>Sign In</h2>
</div>
<div class="sign-in-content">
<form method="post" class="sign-in-form" id="sign-in-form">
<div class="sign-in-form-group">
<label for="username">Username</label>
<input type="text" id="username" name="username" placeholder="Enter Your Username" required>
</div>
<div class="sign-in-form-group">
<label for="password">Password</label>
<input type="password" id="password" name="password" placeholder="Enter Your Password" required>
<span class="material-symbols-outlined visibility">visibility_off</span>
</div>
<div class="sign-in-form-group">
<button type="submit">Sign In</button>
</div>
</form>
</div>
</div>
</div>
`;
    signInCard.innerHTML = signInTemplate;
    return signInCard;
  }
}
