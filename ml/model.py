import pickle
from keras.models import Sequential
from keras.layers import Bidirectional, BatchNormalization
from keras.layers import Dropout, InputLayer, LSTM, CuDNNLSTM
from keras.layers import Dense, BatchNormalization
from tensorflow.keras.optimizers import SGD, Adam, RMSprop

class BILSTM:

    def __init__(self) -> None:
        self.model = self.build()


    def build(self):
        model = Sequential()
        model.add(Bidirectional(LSTM(128, return_sequences=True, input_shape=(WIND_SIZE,6)))) 
        model.add(BatchNormalization())
        model.add(Dropout(0.3)) 
        model.add((LSTM(128, return_sequences=True))) 
        model.add(BatchNormalization())
        model.add(Dropout(0.5))
        model.add(Dense(15, activation="relu"))
        model.add(BatchNormalization())
        model.add(Dropout(0.5))  
        model.add(Dense(4, activation="softmax"))

        return model
    
    def load_model(self, path):
        self.model = pickle.load(open(path, 'rb'))
    
    def save_model(self, path):
        pickle.dump(self.model, open(path, 'wb'))

    def compile(self, model):
        model.compile(loss='categorical_crossentropy',
                    optimizer=Adam(lr=0.0001, decay=1e-6),
                    metrics=['accuracy'])
        return model
    

    def train(self, model, X_train, y_train, X_test, y_test):
        model.fit(X_train, y_train,
                batch_size=64,
                epochs=100,
                validation_data=(X_test, y_test))
        return model
    

    def predict(self, model, X_test):
        return model.predict(X_test)
    
