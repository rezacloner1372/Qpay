const mongoose = require("mongoose");

const connectDB = async() => {
    try {
        const uri = process.env.MONGODB_URI || "";
        await mongoose
            .connect(uri, {
                useNewUrlParser: true,
                authSource: 'admin',
                useUnifiedTopology: true,
            })
            .catch((error) => console.log(error));
        const connection = mongoose.connection;
        console.log("MONGODB CONNECTED SUCCESSFULLY!");
    } catch (error) {
        console.log(error);
        return error;
    }
};

module.exports = connectDB;