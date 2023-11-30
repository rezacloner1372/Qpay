const mongoose = require("mongoose");
const AutoIncrement = require('mongoose-sequence')(mongoose);

const Schema = mongoose.Schema;

const userSchema = Schema({
    fullName: {
        type: String,
    },
    phone: {
        type: String,
    },
    note:{
        type: String,
    },
    createdAt: {
        type: Date,
        default: Date.now,
    }
});
userSchema.plugin(AutoIncrement, {
    inc_field: 'userId',
    start_seq: 1000000
});
module.exports = mongoose.model("User", userSchema);