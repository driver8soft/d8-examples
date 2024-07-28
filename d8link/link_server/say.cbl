      ******************************************************************
      *
      * Loan Calculator Subroutine
      * ==========================
      *
      * A sample program to demonstrate how to create a gRPC COBOL
      * microservice.
      *
      ******************************************************************
       IDENTIFICATION DIVISION.
       PROGRAM-ID. say.
       ENVIRONMENT DIVISION.
       CONFIGURATION SECTION.
       DATA DIVISION.
       FILE SECTION.
       WORKING-STORAGE SECTION.
      * In COBOL, you declare variables in the WORKING-STORAGE section

       LINKAGE SECTION.
      * Data to share with COBOL subroutines 
       01 COMMAREA.
           05 A PIC X(13).
           05 B PIC S9(7)V9(2).
           05 C PIC S9(7)V9(2) COMP-3.
           05 D PIC S9(4) COMP.
           05 E PIC S9(9) COMP.

       PROCEDURE DIVISION USING COMMAREA. 
      * code goes here!
       
           MOVE "Hello, World!" TO A.
           MOVE 12345.67 TO B.
           MOVE -12345.67 TO C.
           MOVE 1234 TO D.
           MOVE 123456789 TO E.


           GOBACK.


