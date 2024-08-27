export type ForgotPasswordType = {
    password: string
    confirm_password: string
}

export type ForgotPasswordResponse = {
    reset_link: string
}

export type ResetPasswordResponse = {
    clicked: boolean
    status: string
}